package router

import (
	"net/http"
	"url-shortener/helper"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.New()

	r.GET("/health", helper.Make(func(c *gin.Context) error {
		c.JSON(http.StatusOK, gin.H{
			"status":     "UP",
			"statusCode": http.StatusOK,
			"message":    "Service is running smoothly",
		})
		return nil
	}))

	// r.POST("/signup", helper.Make(userHandler.SignUpHandler))
}

func Start(addr string) error {
	return r.Run(addr)
}
