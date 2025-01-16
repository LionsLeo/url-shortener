package router

import (
	"net/http"
	"url-shortener/helper"
	"url-shortener/internal/user"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.New()

	r.GET("/health", helper.Make(func(c *gin.Context) error {
		c.JSON(http.StatusOK, gin.H{
			"status":     "UP",
			"statusCode": http.StatusOK,
			"message":    "Service is running smoothly",
		})
		return nil
	}))

	r.POST("/send-otp", helper.Make(userHandler.SendOtpHandler))
	r.POST("/verify-otp", helper.Make(userHandler.VerifyOtp))
}

func Start(addr string) error {
	return r.Run(addr)
}
