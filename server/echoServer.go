package server

import (
	"fmt"
	"leonardodelira/gocleanarch/cockroach/handlers"
	"leonardodelira/gocleanarch/cockroach/repositories"
	"leonardodelira/gocleanarch/cockroach/usecases"
	"leonardodelira/gocleanarch/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type echoServer struct {
	app    *echo.Echo
	db     *gorm.DB
	config *config.Config
}

func NewEchoServer(config *config.Config, db *gorm.DB) Server {
	return &echoServer{
		app:    echo.New(),
		db:     db,
		config: config,
	}
}

func (s *echoServer) Start() {
	s.initializeCockroachHttpHandler()
	s.app.Use(middleware.Logger())
	serverUrl := fmt.Sprintf(":%d", s.config.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeCockroachHttpHandler() {
	// Initialize all layers
	cockroachPostgresRepository := repositories.NewCockroachPostgressRepo(s.db)
	cockroachFCMMessaging := repositories.NewCockroachFCMMessaging()

	cockroachUsecase := usecases.NewCockroachUsecaseImpl(
		cockroachPostgresRepository,
		cockroachFCMMessaging,
	)

	cockroachHttpHandler := handlers.NewCockroachHttp(cockroachUsecase)

	// Routers
	cockroachRouters := s.app.Group("v1/cockroach")
	cockroachRouters.POST("", cockroachHttpHandler.DetectCockroach)
}
