package repositories

import "leonardodelira/gocleanarch/cockroach/entities"

type CockroachRepository interface {
	InsertCockroachData(in *entities.InsertCockroachDto) error
}
