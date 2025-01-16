package user

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type SendOtpRequest struct {
	PhoneNumber string `json:"phone_number"`
	Otp         string `json:"otp"`
}

func (r *SendOtpRequest) Validate() (map[string]string, bool) {
	validationMap := make(map[string]string)
	if r.PhoneNumber == "" {
		validationMap["phone_number"] = "Phone Number is required"
		return validationMap, false
	}

	if len(r.PhoneNumber) != 10 {
		validationMap["phone_number"] = "Invalid Length of Phone Number"
		return validationMap, false
	}

	if len(validationMap) == 0 {
		return validationMap, true
	} else {
		return validationMap, false
	}
}

func (r *SendOtpRequest) SetOtp(otp string) {
	r.Otp = otp
}

type SendOtpResponse struct {
	PhoneNumber string `json:"phone_number"`
	AccessToken string `json:"access_token"`
}

func NewSendOtpResponse(res *SendOtpRequest, accessToken string) *SendOtpResponse {
	return &SendOtpResponse{
		PhoneNumber: res.PhoneNumber,
		AccessToken: accessToken,
	}
}

// MyCustomClaims defines your custom JWT claims structure.
// In addition to custom fields, we embed jwt.RegisteredClaims
// to include standard fields like exp, iat, and iss.
type MyCustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// secretKey is used to sign tokens (HMAC).
// In a real application, store this securely (env var, vault, etc.).
var secretKey = os.Getenv("AUTH_SECRET")

// secretKey is used to sign tokens (HMAC).
// In a real application, store this securely (env var, vault, etc.).
var secretKeyForApplication = os.Getenv("APP_SECRET")

type VerifyOtpRequest struct {
	Otp string `json:"otp"`
}

type VerifyOtpResponse struct {
	AccessToken string `json:"access_token"`
}
