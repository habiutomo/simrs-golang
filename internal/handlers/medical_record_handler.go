package handlers

import (
	"strconv"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/services"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type MedicalRecordHandler struct {
	service *services.MedicalRecordService
}

func NewMedicalRecordHandler(service *services.MedicalRecordService) *MedicalRecordHandler {
	return &MedicalRecordHandler{service: service}
}

func (h *MedicalRecordHandler) Create(c *gin.Context) {
	var req dto.CreateMedicalRecordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	record, err := h.service.Create(&req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Medical record created", record)
}

func (h *MedicalRecordHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	record, err := h.service.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Medical record not found")
		return
	}

	response.OK(c, "Medical record retrieved", record)
}

func (h *MedicalRecordHandler) GetByPatient(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patientId"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid patient ID", nil)
		return
	}

	records, err := h.service.GetByPatient(uint(patientID))
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve records")
		return
	}

	response.OK(c, "Medical records retrieved", records)
}

func (h *MedicalRecordHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	records, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve records")
		return
	}

	response.Paginated(c, "Medical records retrieved", records, page, limit, int(total))
}

func (h *MedicalRecordHandler) AddPrescription(c *gin.Context) {
	recordID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid record ID", nil)
		return
	}

	var req dto.CreatePrescriptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	prescription, err := h.service.AddPrescription(uint(recordID), &req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Prescription added", prescription)
}
