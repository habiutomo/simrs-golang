package handlers

import (
	"simrs-golang/internal/repositories"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type DepartmentHandler struct {
	repo *repositories.DepartmentRepository
}

func NewDepartmentHandler(repo *repositories.DepartmentRepository) *DepartmentHandler {
	return &DepartmentHandler{repo: repo}
}

func (h *DepartmentHandler) GetAll(c *gin.Context) {
	departments, err := h.repo.FindAll()
	if err != nil {
		response.InternalServerError(c, "Failed to retrieve departments")
		return
	}

	response.OK(c, "Departments retrieved", departments)
}
