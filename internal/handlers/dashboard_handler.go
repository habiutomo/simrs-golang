package handlers

import (
	"simrs-golang/internal/services"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service *services.DashboardService
}

func NewDashboardHandler(service *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: service}
}

func (h *DashboardHandler) GetDashboard(c *gin.Context) {
	dashboard, err := h.service.GetDashboard()
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve dashboard")
		return
	}

	response.OK(c, "Dashboard data retrieved", dashboard)
}
