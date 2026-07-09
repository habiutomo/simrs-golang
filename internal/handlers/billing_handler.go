package handlers

import (
	"strconv"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/services"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type BillingHandler struct {
	service *services.BillingService
}

func NewBillingHandler(service *services.BillingService) *BillingHandler {
	return &BillingHandler{service: service}
}

func (h *BillingHandler) Create(c *gin.Context) {
	var req dto.CreateBillingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	billing, err := h.service.Create(&req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.Created(c, "Billing created", billing)
}

func (h *BillingHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	billing, err := h.service.GetByID(uint(id))
	if err != nil {
		response.NotFound(c, "Billing not found")
		return
	}

	response.OK(c, "Billing retrieved", billing)
}

func (h *BillingHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	status := c.Query("status")

	billings, total, err := h.service.GetAll(page, limit, status)
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve billings")
		return
	}

	response.Paginated(c, "Billings retrieved", billings, page, limit, int(total))
}

func (h *BillingHandler) GetByPatient(c *gin.Context) {
	patientID, err := strconv.ParseUint(c.Param("patientId"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid patient ID", nil)
		return
	}

	billings, err := h.service.GetByPatient(uint(patientID))
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve billings")
		return
	}

	response.OK(c, "Billings retrieved", billings)
}

func (h *BillingHandler) Pay(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid ID", nil)
		return
	}

	var req dto.PayBillingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	billing, err := h.service.Pay(uint(id), &req)
	if err != nil {
		response.BadRequest(c, err.Error(), nil)
		return
	}

	response.OK(c, "Payment successful", billing)
}
