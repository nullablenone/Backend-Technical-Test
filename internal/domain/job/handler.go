package job

import (
	"net/http"
	"redikru-test/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{Service: service}
}



// @Summary      Membuat lowongan pekerjaan baru
// @Description  Menambahkan data lowongan pekerjaan baru ke dalam sistem
// @Tags         Lowongan Pekerjaan
// @Accept       json
// @Produce      json
// @Param        payload body CreateJobRequest true "Payload untuk membuat lowongan baru"
// @Success      201 {object} ResponseSuccessJob "Sukses membuat lowongan"
// @Failure      400 {object} utils.ResponseError "Request payload tidak valid"
// @Failure      500 {object} utils.ResponseError "Terjadi kesalahan internal pada server"
// @Router       /jobs [post]
func (h *Handler) CreateJobHandler(c *gin.Context) {
	var request CreateJobRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	newJob, err := h.Service.CreateJob(request)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondSuccess(c, http.StatusCreated, newJob, "Job posting created successfully")

}


// @Summary      Mendapatkan daftar lowongan pekerjaan
// @Description  Mengambil daftar lowongan pekerjaan dengan urutan data terbaru dan dilengkapi dengan filter dan pagination.
// @Tags         Lowongan Pekerjaan
// @Produce      json
// @Param        keyword query string false "Pencarian berdasarkan kata kunci pada judul dan deskripsi lowongan."
// @Param        companyName query string false "Filter berdasarkan nama perusahaan."
// @Param        page query int false "Nomor halaman yang ingin ditampilkan." default(1)
// @Param        limit query int false "Jumlah data yang ditampilkan per halaman." default(10)
// @Success      200 {object} ResponseSuccessGetJobs "Sukses mengambil data lowongan"
// @Failure      400 {object} utils.ResponseError "Parameter query tidak valid"
// @Failure      500 {object} utils.ResponseError "Terjadi kesalahan internal pada server"
// @Router       /jobs [get]
func (h *Handler) GetAllJobHandler(c *gin.Context) {
	var request GetAllJobsRequest

	err := c.ShouldBindQuery(&request)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Parameter query tidak valid")
		return
	}

	if request.Page <= 0 {
		request.Page = 1
	}
	if request.Limit <= 0 {
		request.Limit = 10
	}

	if request.Limit > 100 {
		request.Limit = 100
	}

	jobs, pagination, err := h.Service.GetAllJob(request)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.RespondSuccessWithPagination(c, http.StatusOK, jobs, pagination, "Permintaan berhasil")

}
