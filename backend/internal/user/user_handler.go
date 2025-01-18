package user

import (
	"fmt"
	"net/http"
	"url-shortener/helper"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Handler struct {
	UserService *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{
		UserService: s,
	}
}

func (h *Handler) SendOtpHandler(c *gin.Context) error {
	var u SendOtpRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		return helper.InvalidJson()
	}
	if validationMap, ok := u.Validate(); !ok {
		return helper.NewApiError(http.StatusBadRequest, fmt.Errorf("validation error"), validationMap)
	}

	res, err := h.UserService.SendOtpService(&u, c)

	if err != nil {
		return err
	}

	helper.WriteJson(c, http.StatusCreated, res)
	return nil
}

func (h *Handler) VerifyOtp(c *gin.Context) error {
	token, err := helper.GetToken(c)
	if err != nil {
		return err
	}

	claims, err := validateJwt(token)
	if err != nil {
		return err
	}

	if claims.Role != "auth-verify-otp" {
		return helper.NewApiError(http.StatusUnauthorized, fmt.Errorf("invalid Token"), "")
	}

	var u VerifyOtpRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		return helper.InvalidJson()
	}

	res, err := h.UserService.VerifyOtpService(&u, c, claims)

	if err != nil {
		return err
	}

	helper.WriteJson(c, http.StatusCreated, res)
	return nil
}

// ValidateJWT parses the token string, verifies the signature,
// and returns the custom claims if valid.
func validateJwt(tokenString string) (*MyCustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token’s signing method matches our expectation (HMAC with SHA-256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	// Validate token and return claims
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, helper.NewApiError(http.StatusUnauthorized, fmt.Errorf("token is invalid"), "")
}
