package server

import (
	"fmt"
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

func NewEchoServer(config config.Config, db *gorm.DB) Server {
	return &echoServer{
		app:    echo.New(),
		db:     db,
		config: &config,
	}
}

func (s *echoServer) Start() {

	s.app.Use(middleware.Logger())
	serverUrl := fmt.Sprintf(":%d", s.config.App.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
