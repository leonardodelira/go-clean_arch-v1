package main

import (
	"leonardodelira/gocleanarch/config"
	"leonardodelira/gocleanarch/database"
	"leonardodelira/gocleanarch/server"
)

func main() {
	config := config.GetConfig()
	db := database.NewPostgresDatabase(&config)
	server.NewEchoServer(&config, db.GetDb()).Start()
}
