package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response adalah struktur standar response API sesuai API Specification
type Response struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Errors     []Error     `json:"errors,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	RetryAfter int         `json:"retry_after,omitempty"`
}

// Error adalah struktur error per field
type Error struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

// Pagination adalah struktur pagination
type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
	HasNext    bool  `json:"has_next"`
	HasPrev    bool  `json:"has_prev"`
}

// Success mengirim response sukses (200)
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created mengirim response sukses untuk resource baru (201)
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// NoContent mengirim response 204 (no content)
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// ErrorResponse mengirim response error dengan status code
func ErrorResponse(c *gin.Context, statusCode int, message string, errors []Error) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// BadRequest mengirim response 400
func BadRequest(c *gin.Context, message string, errors []Error) {
	ErrorResponse(c, http.StatusBadRequest, message, errors)
}

// Unauthorized mengirim response 401
func Unauthorized(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message, nil)
}

// Forbidden mengirim response 403
func Forbidden(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusForbidden, message, nil)
}

// NotFound mengirim response 404
func NotFound(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message, nil)
}

// Conflict mengirim response 409
func Conflict(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusConflict, message, nil)
}

// UnprocessableEntity mengirim response 422 (validation error)
func UnprocessableEntity(c *gin.Context, message string, errors []Error) {
	ErrorResponse(c, http.StatusUnprocessableEntity, message, errors)
}

// TooManyRequests mengirim response 429 (rate limit exceeded)
func TooManyRequests(c *gin.Context, message string, retryAfter int) {
	c.JSON(http.StatusTooManyRequests, Response{
		Success:    false,
		Message:    message,
		Errors:     []Error{},
		RetryAfter: retryAfter,
	})
}

// InternalServerError mengirim response 500
func InternalServerError(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message, nil)
}

// Paginated mengirim response dengan pagination
func Paginated(c *gin.Context, message string, data interface{}, pagination Pagination) {
	c.JSON(http.StatusOK, Response{
		Success:    true,
		Message:    message,
		Data:       data,
		Pagination: &pagination,
	})
}

// NewPagination membuat pagination struct dari parameter
func NewPagination(page, limit int, total int64) Pagination {
	totalPages := int(total) / limit
	if int(total)%limit > 0 {
		totalPages++
	}

	return Pagination{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
}