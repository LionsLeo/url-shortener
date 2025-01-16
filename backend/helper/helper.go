package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"

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

func CreateMd5Hash(text string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", NewApiError(http.StatusUnauthorized, fmt.Errorf("Authorization header missing"), "")
	}

	// Typically tokens are passed as: "Authorization: Bearer <token>"
	// So we split and check if it starts with "Bearer "
	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		return "", NewApiError(http.StatusBadRequest, fmt.Errorf("Authorization header format must be Bearer <token>"), "")
	}

	// Extract the token by removing "Bearer "
	token := strings.TrimPrefix(authHeader, bearerPrefix)
	return token, nil
}
