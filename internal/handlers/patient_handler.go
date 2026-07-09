package handlers

import (
	"strconv"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/services"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	service *services.PatientService
}

func NewPatientHandler(service *services.PatientService) *PatientHandler {
	return &PatientHandler{service: service}
}

func (h *PatientHandler) Create(c *gin.Context) {
	var req dto.CreatePatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	patient, err := h.service.Create(&req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Patient created successfully", patient)
}

func (h *PatientHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	patient, err := h.service.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Patient not found")
		return
	}

	response.OK(c, "Patient retrieved", patient)
}

func (h *PatientHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	patients, total, err := h.service.GetAll(page, limit, search)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve patients")
		return
	}

	response.Paginated(c, "Patients retrieved", patients, page, limit, int(total))
}

func (h *PatientHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	var req dto.UpdatePatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	patient, err := h.service.Update(uint(id), &req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.OK(c, "Patient updated", patient)
}

func (h *PatientHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		response.InternalServerError(c, "Failed to delete patient")
		return
	}

	response.OK(c, "Patient deleted", nil)
}
