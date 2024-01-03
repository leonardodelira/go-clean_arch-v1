package repositories

import "leonardodelira/gocleanarch/cockroach/entities"

type CockroachMessaging interface {
	PushNotification(m *entities.CockroachPushNotificationDto) error
}
