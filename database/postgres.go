package database

import (
	"fmt"
	"leonardodelira/gocleanarch/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(config *config.Config) Database {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Db.Host,
		config.Db.User,
		config.Db.Password,
		config.Db.DBName,
		config.Db.Port,
		config.Db.SSLMode,
		config.Db.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDatabase{Db: db}
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return p.Db
}
