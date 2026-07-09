package handlers

import (
	"strconv"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/services"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type MedicationHandler struct {
	service *services.MedicationService
}

func NewMedicationHandler(service *services.MedicationService) *MedicationHandler {
	return &MedicationHandler{service: service}
}

func (h *MedicationHandler) Create(c *gin.Context) {
	var req dto.CreateMedicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	medication, err := h.service.Create(&req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Medication created", medication)
}

func (h *MedicationHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	medication, err := h.service.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Medication not found")
		return
	}

	response.OK(c, "Medication retrieved", medication)
}

func (h *MedicationHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	medications, total, err := h.service.GetAll(page, limit, search)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve medications")
		return
	}

	response.Paginated(c, "Medications retrieved", medications, page, limit, int(total))
}

func (h *MedicationHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	var req dto.CreateMedicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	medication, err := h.service.Update(uint(id), &req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.OK(c, "Medication updated", medication)
}

func (h *MedicationHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		response.InternalServerError(c, "Failed to delete medication")
		return
	}

	response.OK(c, "Medication deleted", nil)
}
