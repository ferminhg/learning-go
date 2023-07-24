package main

import (
	"github.com/ferminhg/learning-go/internal/infra/routing"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupServer()
	r.Run(":8080")
}

func setupServer() *gin.Engine {
	r := gin.Default()
	controller := routing.New()

	r.GET("/health", controller.GetHealthEndpoint)

	r.PUT("/ads", controller.PostNewAdsEndpoint)
	r.POST("/ads", controller.PostNewAdsEndpoint)
	r.GET("/ads/:id", controller.GetAdByIdEndpoint)
	r.GET("/ads", controller.GetAdsEndpoint)

	return r
}
