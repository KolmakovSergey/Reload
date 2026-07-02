package repository

import "reload/internal/models"

type EventStorage interface {
	SaveEvent(event models.Event) error
	GetEventsByUserId(id int) ([]models.Event, error)
	GetAllEvents() ([]models.Event, error)
}