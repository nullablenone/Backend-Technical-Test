package job

import (
	"encoding/json"
	"redikru-test/utils"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2" 
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mock repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateJob(job *Job) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *MockRepository) GetAllJob(request GetAllJobsRequest) ([]Job, int64, error) {
	args := m.Called(request)
	return args.Get(0).([]Job), args.Get(1).(int64), args.Error(2)
}


// test function CreateJob success
func TestCreateJob_Success(t *testing.T) {
	mr, err := miniredis.Run()
	assert.NoError(t, err, "Gagal menjalankan miniredis")
	defer mr.Close() 

	redisClient := redis.NewClient(&redis.Options{
		Addr: mr.Addr(), 
	})

	mockRepo := new(MockRepository)
	jobService := NewService(mockRepo, redisClient)

	request := CreateJobRequest{
		Title:       "Software Engineer",
		Description: "Develop amazing software.",
		CompanyID:   "company-123",
	}

	mr.Set("jobs:lama", "data-lama")
	mr.Set("jobs:lain", "data-lain")
	assert.True(t, mr.Exists("jobs:lama")) 

	mockRepo.On("CreateJob", mock.AnythingOfType("*job.Job")).Return(nil)

	_, err = jobService.CreateJob(request)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	assert.False(t, mr.Exists("jobs:lama"), "Cache 'jobs:lama' seharusnya sudah terhapus")
	assert.False(t, mr.Exists("jobs:lain"), "Cache 'jobs:lain' seharusnya sudah terhapus")
}

// test function GetAllJob ketika cache nya miss atau tidak ada
func TestGetAllJob_CacheMiss(t *testing.T) {
	mr, err := miniredis.Run()
	assert.NoError(t, err)
	defer mr.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockRepo := new(MockRepository)
	jobService := NewService(mockRepo, redisClient)

	request := GetAllJobsRequest{Page: 1, Limit: 5, Keyword: "Backend"}
	cacheKey := "jobs:Backend::1:5"
	mockJobs := []Job{{ID: "job-1", Title: "Backend Engineer (dari DB)"}}
	var totalRecords int64 = 1

	mockRepo.On("GetAllJob", request).Return(mockJobs, totalRecords, nil).Once()

	jobs, pagination, err := jobService.GetAllJob(request)

	assert.NoError(t, err)
	assert.Len(t, jobs, 1, "Jumlah jobs harus 1")
	assert.Equal(t, "Backend Engineer (dari DB)", jobs[0].Title)
	assert.Equal(t, totalRecords, pagination.TotalRecords)
	assert.Equal(t, 1, pagination.TotalPages) 

	mockRepo.AssertExpectations(t)

	assert.True(t, mr.Exists(cacheKey), "Cache seharusnya sudah dibuat setelah cache miss")
}

// test function GetAllJob ketika cache nya di Hit atau ada
func TestGetAllJob_CacheHit(t *testing.T) {
	mr, err := miniredis.Run()
	assert.NoError(t, err)
	defer mr.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	mockRepo := new(MockRepository)
	jobService := NewService(mockRepo, redisClient)

	request := GetAllJobsRequest{Page: 1, Limit: 5, Keyword: "Backend"}
	cacheKey := "jobs:Backend::1:5"

	cachedResponse := cachedJobsResponse{
		Jobs:       []Job{{ID: "job-cached", Title: "Backend Engineer (dari Cache)"}},
		Pagination: utils.Pagination{CurrentPage: 1, PerPage: 5, TotalPages: 1, TotalRecords: 1},
	}
	cachedJSON, _ := json.Marshal(cachedResponse)

	mr.Set(cacheKey, string(cachedJSON))
	mr.SetTTL(cacheKey, 10*time.Minute)


	jobs, pagination, err := jobService.GetAllJob(request)

	assert.NoError(t, err)
	assert.Len(t, jobs, 1)
	assert.Equal(t, "Backend Engineer (dari Cache)", jobs[0].Title, "Data harus berasal dari cache")
	assert.Equal(t, cachedResponse.Pagination.TotalRecords, pagination.TotalRecords)

	mockRepo.AssertNotCalled(t, "GetAllJob", request)
}
