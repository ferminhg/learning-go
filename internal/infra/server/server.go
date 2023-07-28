package server

import (
	"fmt"
	"github.com/ferminhg/learning-go/internal/application"
	"github.com/ferminhg/learning-go/internal/infra/handler"
	"github.com/ferminhg/learning-go/internal/infra/storage/inmemory"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
}

func New(host string, port uint) Server {
	srv := Server{
		engine:   gin.Default(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) registerRoutes() {
	service := application.NewAdService(inmemory.NewInMemoryAdRepository())

	s.engine.GET("/health", handler.GetHealthEndpoint())
	s.engine.PUT("/ads", handler.PostNewAdsEndpoint(service))
	s.engine.POST("/ads", handler.PostNewAdsEndpoint(service))
	s.engine.GET("/ads/:id", handler.GetAdByIdEndpoint(service))
	s.engine.GET("/ads", handler.GetAdsEndpoint(service))
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}
