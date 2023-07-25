package server

import (
	"fmt"
	"github.com/ferminhg/learning-go/internal/infra/routing"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
}

func (s Server) registerRoutes() {
	controller := routing.New()

	s.engine.GET("/health", controller.GetHealthEndpoint)

	s.engine.PUT("/ads", controller.PostNewAdsEndpoint)
	s.engine.POST("/ads", controller.PostNewAdsEndpoint)
	s.engine.GET("/ads/:id", controller.GetAdByIdEndpoint)
	s.engine.GET("/ads", controller.GetAdsEndpoint)
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.Default(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}
