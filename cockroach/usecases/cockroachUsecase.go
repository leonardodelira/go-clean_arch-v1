package usecases

import (
	"leonardodelira/gocleanarch/cockroach/entities"
	"leonardodelira/gocleanarch/cockroach/models"
	"leonardodelira/gocleanarch/cockroach/repositories"
	"time"
)

type CockroachUsecase interface {
	CockroachDataProcessing(in *models.AddCockroachData) error
}

type cockroachUsecaseImpl struct {
	cockroachRepository repositories.CockroachRepository
	cockroachMessaging  repositories.CockroachMessaging
}

func NewCockroachUsecaseImpl(repo repositories.CockroachRepository, msg repositories.CockroachMessaging) CockroachUsecase {
	return &cockroachUsecaseImpl{
		cockroachRepository: repo,
		cockroachMessaging:  msg,
	}
}

func (u *cockroachUsecaseImpl) CockroachDataProcessing(in *models.AddCockroachData) error {
	insertCockroachData := &entities.InsertCockroachDto{
		Amount: in.Amount,
	}
	if err := u.cockroachRepository.InsertCockroachData(insertCockroachData); err != nil {
		return err
	}

	notificationCockroachData := &entities.CockroachPushNotificationDto{
		Title:        "Cockroach Detected 🪳 !!!",
		Amount:       in.Amount,
		ReportedTime: time.Now().Local().Format("2006-01-02 15:04:05"),
	}
	if err := u.cockroachMessaging.PushNotification(notificationCockroachData); err != nil {
		return err
	}

	return nil
}
