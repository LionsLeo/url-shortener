package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
	"url-shortener/helper"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/rand"
)

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) SendOtpService(req *SendOtpRequest, ctx *gin.Context) (*SendOtpResponse, error) {
	otp := generateOTP(4)
	if os.Getenv("APP_MODE") == "dev" {
		otp = "1234"
	}
	req.SetOtp(otp)
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshalling to JSON: %v", err)
	}
	hashKey := helper.CreateMd5Hash(req.PhoneNumber)
	key := "users-otp:" + hashKey
	if err := s.Repo.CreateOtpEntry(key, jsonData, ctx); err != nil {
		return nil, fmt.Errorf("error storing in redis: %v", err)
	}
	to := "+91" + req.PhoneNumber
	err = s.Repo.SendSmsOtp(to, otp)
	if err != nil {
		return nil, err
	}
	token, err := generateJWTForVerifyOtp(hashKey)
	if err != nil {
		return nil, err
	}

	return NewSendOtpResponse(req, token), nil
}

func (s *Service) VerifyOtpService(req *VerifyOtpRequest, ctx *gin.Context, claims *MyCustomClaims) (*VerifyOtpResponse, error) {
	key := "users-otp:" + claims.Username
	otp, err := s.Repo.GetOtp(key, ctx)
	if err != nil {
		return nil, err
	}

	if otp != req.Otp {
		return nil, helper.NewApiError(http.StatusUnauthorized, fmt.Errorf("OTP mismatch"), "")
	}
	// TODO create a user row in the postgres table and remove the otp key from the redis and send the generated username for jwt generation.
	token, err := generateJWTForApplication("my")
	if err != nil {
		return nil, err
	}
	res := VerifyOtpResponse{AccessToken: token}

	return &res, nil
}

// generateOTP creates a random numeric OTP of the specified length.
func generateOTP(length int) string {
	rand.Seed(uint64(time.Now().UnixNano()))
	digits := "0123456789"
	otpBytes := make([]byte, length)
	for i := 0; i < length; i++ {
		otpBytes[i] = digits[rand.Intn(len(digits))]
	}
	return string(otpBytes)
}

func generateJWTForVerifyOtp(username string) (string, error) {
	// Create the Claims
	claims := MyCustomClaims{
		Username: username,
		Role:     "auth-verify-otp",
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt is a pointer, so we use jwt.NewNumericDate for convenience
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)), // Token valid for 5 min
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app",
		},
	}

	// Create token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateJWTForApplication(username string) (string, error) {
	// Create the Claims
	claims := MyCustomClaims{
		Username: username,
		Role:     "application",
		RegisteredClaims: jwt.RegisteredClaims{
			// ExpiresAt is a pointer, so we use jwt.NewNumericDate for convenience
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), // Token valid for 5 min
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app",
		},
	}

	// Create token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(secretKeyForApplication))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
