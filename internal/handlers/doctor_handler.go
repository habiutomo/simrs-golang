package handlers

import (
	"strconv"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/services"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type DoctorHandler struct {
	service *services.DoctorService
}

func NewDoctorHandler(service *services.DoctorService) *DoctorHandler {
	return &DoctorHandler{service: service}
}

func (h *DoctorHandler) Create(c *gin.Context) {
	var req dto.CreateDoctorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	doctor, err := h.service.Create(&req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Doctor created successfully", doctor)
}

func (h *DoctorHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	doctor, err := h.service.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Doctor not found")
		return
	}

	response.OK(c, "Doctor retrieved", doctor)
}

func (h *DoctorHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	doctors, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve doctors")
		return
	}

	response.Paginated(c, "Doctors retrieved", doctors, page, limit, int(total))
}

func (h *DoctorHandler) GetByDepartment(c *gin.Context) {
	deptID, err := strconv.ParseUint(c.Param("deptId"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid department ID", nil)
		return
	}

	doctors, err := h.service.GetByDepartment(uint(deptID))
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve doctors")
		return
	}

	response.OK(c, "Doctors retrieved", doctors)
}

func (h *DoctorHandler) GetAvailable(c *gin.Context) {
	doctors, err := h.service.GetAllAvailable()
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve doctors")
		return
	}

	response.OK(c, "Available doctors retrieved", doctors)
}

func (h *DoctorHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		response.InternalServerError(c, "Failed to delete doctor")
		return
	}

	response.OK(c, "Doctor deleted", nil)
}

func (h *DoctorHandler) AddSchedule(c *gin.Context) {
	var req dto.UpdateDoctorScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	schedule, err := h.service.AddSchedule(&req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Schedule added", schedule)
}

func (h *DoctorHandler) GetSchedules(c *gin.Context) {
	doctorID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid doctor ID", nil)
		return
	}

	schedules, err := h.service.GetSchedules(uint(doctorID))
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve schedules")
		return
	}

	response.OK(c, "Schedules retrieved", schedules)
}
