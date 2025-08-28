package job

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"redikru-test/utils"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Service interface {
	CreateJob(request CreateJobRequest) (*Job, error)
	GetAllJob(request GetAllJobsRequest) ([]Job, utils.Pagination, error)
}

type service struct {
	repository Repository
	redis      *redis.Client
	ctx        context.Context
}

func NewService(repo Repository, redisClient *redis.Client) Service {
	return &service{repository: repo, redis: redisClient, ctx: context.Background()}
}

func (s *service) CreateJob(request CreateJobRequest) (*Job, error) {
	job := Job{
		ID:          uuid.NewString(),
		Title:       request.Title,
		Description: request.Description,
		CompanyID:   request.CompanyID,
	}

	err := s.repository.CreateJob(&job)
	if err != nil {
		return nil, err
	}

	iter := s.redis.Scan(s.ctx, 0, "jobs:*", 0).Iterator()
	for iter.Next(s.ctx) {
		s.redis.Del(s.ctx, iter.Val())
	}
	if err := iter.Err(); err != nil {
		log.Printf("Peringatan: Gagal menghapus cache Redis: %v", err)
	}

	return &job, nil
}

func (s *service) GetAllJob(request GetAllJobsRequest) ([]Job, utils.Pagination, error) {

	// key redis
	cacheKey := fmt.Sprintf("jobs:%s:%s:%d:%d", request.Keyword, request.CompanyName, request.Page, request.Limit)

	// cek cache
	cachedData, err := s.redis.Get(s.ctx, cacheKey).Result()
	if err == nil {
		log.Println("CACHE HIT!")
		var response cachedJobsResponse
		if err := json.Unmarshal([]byte(cachedData), &response); err == nil {
			return response.Jobs, response.Pagination, nil
		}
	}

	log.Println("CACHE MISS! Mengambil dari database...")
	jobs, totalRecords, err := s.repository.GetAllJob(request)
	if err != nil {
		return nil, utils.Pagination{}, err
	}

	// total halaman
	totalPages := int(math.Ceil(float64(totalRecords) / float64(request.Limit)))

	// masukan data pagination
	pagination := utils.Pagination{
		CurrentPage:  request.Page,
		PerPage:      request.Limit,
		TotalPages:   totalPages,
		TotalRecords: totalRecords,
	}

	// set cache
	responseToCache := cachedJobsResponse{
		Jobs:       jobs,
		Pagination: pagination,
	}

	jsonData, err := json.Marshal(responseToCache)
	if err != nil {
		log.Printf("Peringatan: Gagal mengubah data menjadi JSON untuk cache: %v", err)
	} else {
		s.redis.Set(s.ctx, cacheKey, jsonData, 10*time.Minute)
	}

	return jobs, pagination, nil
}
