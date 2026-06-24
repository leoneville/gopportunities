package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
