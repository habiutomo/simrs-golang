package handlers

import (
	"strconv"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/services"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	service *services.AppointmentService
}

func NewAppointmentHandler(service *services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

func (h *AppointmentHandler) Create(c *gin.Context) {
	var req dto.CreateAppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	appointment, err := h.service.Create(&req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Appointment created", appointment)
}

func (h *AppointmentHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	appointment, err := h.service.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Appointment not found")
		return
	}

	response.OK(c, "Appointment retrieved", appointment)
}

func (h *AppointmentHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	status := c.Query("status")

	appointments, total, err := h.service.GetAll(page, limit, status)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve appointments")
		return
	}

	response.Paginated(c, "Appointments retrieved", appointments, page, limit, int(total))
}

func (h *AppointmentHandler) GetByPatient(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patientId"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid patient ID", nil)
		return
	}

	appointments, err := h.service.GetByPatient(uint(patientID))
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve appointments")
		return
	}

	response.OK(c, "Appointments retrieved", appointments)
}

func (h *AppointmentHandler) GetByDoctor(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("doctorId"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid doctor ID", nil)
		return
	}
	date := c.Query("date")

	appointments, err := h.service.GetByDoctor(uint(doctorID), date)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve appointments")
		return
	}

	response.OK(c, "Appointments retrieved", appointments)
}

func (h *AppointmentHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	appointment, err := h.service.UpdateStatus(uint(id), req.Status)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.OK(c, "Appointment status updated", appointment)
}
