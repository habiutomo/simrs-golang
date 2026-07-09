package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type PaginationMeta struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalRows int `json:"total_rows"`
	TotalPage int `json:"total_page"`
}

type PaginatedResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    interface{}     `json:"data"`
	Meta    PaginationMeta  `json:"meta"`
}

func OK(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string, errors interface{}) {
	c.JSON(status, Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

func BadRequest(c *gin.Context, message string, errors interface{}) {
	Error(c, http.StatusBadRequest, message, errors)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, message, nil)
}

func InternalServerError(c *gin.Context, message string) {
	Error(c, http.StatusInternalServerError, message, nil)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, message, nil)
}

func Paginated(c *gin.Context, message string, data interface{}, page, limit, totalRows int) {
	totalPage := (totalRows + limit - 1) / limit
	c.JSON(http.StatusOK, PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta: PaginationMeta{
			Page:      page,
			Limit:     limit,
			TotalRows: totalRows,
			TotalPage: totalPage,
		},
	})
}
