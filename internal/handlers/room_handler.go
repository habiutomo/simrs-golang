package handlers

import (
	"strconv"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/services"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type RoomHandler struct {
	service *services.RoomService
}

func NewRoomHandler(service *services.RoomService) *RoomHandler {
	return &RoomHandler{service: service}
}

func (h *RoomHandler) GetAll(c *gin.Context) {
	rooms, err := h.service.GetAll()
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve rooms")
		return
	}

	response.OK(c, "Rooms retrieved", rooms)
}

func (h *RoomHandler) GetAvailable(c *gin.Context) {
	rooms, err := h.service.GetAvailable()
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve rooms")
		return
	}

	response.OK(c, "Available rooms retrieved", rooms)
}

func (h *RoomHandler) AdmitPatient(c *gin.Context) {
	var req dto.AdmitPatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	inpatient, err := h.service.AdmitPatient(&req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Patient admitted", inpatient)
}

func (h *RoomHandler) DischargePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	inpatient, err := h.service.DischargePatient(uint(id))
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.OK(c, "Patient discharged", inpatient)
}

func (h *RoomHandler) GetActiveInpatients(c *gin.Context) {
	inpatients, err := h.service.GetActiveInpatients()
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve inpatients")
		return
	}

	response.OK(c, "Active inpatients retrieved", inpatients)
}

func (h *RoomHandler) GetInpatientByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	inpatient, err := h.service.GetInpatientByID(uint(id))
	if err != nil {
		response.NotFound(c, "Inpatient not found")
		return
	}

	response.OK(c, "Inpatient retrieved", inpatient)
}
