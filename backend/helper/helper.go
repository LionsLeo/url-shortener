package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	StatusCode int `json:"statusCode"`
	Msg        any `json:"message"`
	Path       any `json:"path"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%d", e.Msg)
}

func NewApiError(statusCode int, err error, path any) ApiError {
	return ApiError{StatusCode: statusCode, Msg: err.Error(), Path: path}
}

func InvalidJson() ApiError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("invalid JSON request data"), make(map[string]string))
}

type ApiFunc func(c *gin.Context) error

func Make(h ApiFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			if apiErr, ok := err.(ApiError); ok {
				WriteJson(c, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"message":    "internal server error",
				}
				WriteJson(c, http.StatusInternalServerError, errResp)
			}
		}
	}
}

func WriteJson(c *gin.Context, status int, v any) {
	c.Header("Content-Type", "application/json")
	c.JSON(status, v)
}