package repositories

import (
	"leonardodelira/gocleanarch/cockroach/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type cockroachPostgressRepo struct {
	db *gorm.DB
}

func NewCockroachPostgressRepo(db *gorm.DB) CockroachRepository {
	return &cockroachPostgressRepo{db: db}
}

func (r *cockroachPostgressRepo) InsertCockroachData(in *entities.InsertCockroachDto) error {
	data := &entities.Cockroach{
		Amount: in.Amount,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		log.Errorf("InsertCockroachData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertCockroachData: %v", result.RowsAffected)
	return nil
}
