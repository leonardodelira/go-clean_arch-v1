package main

import (
	"leonardodelira/gocleanarch/cockroach/entities"
	"leonardodelira/gocleanarch/config"
	"leonardodelira/gocleanarch/database"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)

	cockroachMigrate(db)
}

func cockroachMigrate(db database.Database) {
	db.GetDb().Migrator().CreateTable(&entities.Cockroach{})
	db.GetDb().CreateInBatches([]entities.Cockroach{
		{Amount: 1},
		{Amount: 2},
		{Amount: 2},
		{Amount: 5},
		{Amount: 3},
	}, 10)
}
