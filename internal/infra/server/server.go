package server

import (
	"fmt"
	"github.com/ferminhg/learning-go/internal/application"
	"github.com/ferminhg/learning-go/internal/infra/eventHandler"
	"github.com/ferminhg/learning-go/internal/infra/generator"
	"github.com/ferminhg/learning-go/internal/infra/handler"
	"github.com/ferminhg/learning-go/internal/infra/storage/inmemory"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	brokerList []string
}

func New(host string, port uint, brokerList []string) Server {
	srv := Server{
		engine:     gin.Default(),
		httpAddr:   fmt.Sprintf("%s:%d", host, port),
		brokerList: brokerList,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) registerRoutes() {
	service := application.NewAdService(
		inmemory.NewInMemoryAdRepository(),
		generator.New(true),
		eventHandler.NewKafkaEventHandler(s.brokerList),
	)

	s.engine.GET("/health", handler.GetHealthEndpoint())

	s.engine.PUT("/ads", handler.PostNewAdsEndpoint(service))
	s.engine.POST("/ads", handler.PostNewAdsEndpoint(service))

	s.engine.GET("/ads/:id", handler.GetAdByIdEndpoint(service))
	s.engine.GET("/ads", handler.GetAdsEndpoint(service))

	s.engine.DELETE("/ads/:id", handler.DeleteAdByIdHandler(service))

	s.engine.POST("/description-generator", handler.PostDescriptionGenerator(service))
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}
