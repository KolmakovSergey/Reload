package repository

import "reload/internal/models"

type EventRepo struct {
	Storage []models.Event
}

func NewEventRepo(events []models.Event) *EventRepo {
	return &EventRepo{
		Storage: events,
	}
}

func (e *EventRepo) SaveEvent(event models.Event) {
	e.Storage = append(e.Storage, event)
}

func (e *EventRepo) GetEventsByUserId(id int) []models.Event {

	var tempSlice []models.Event

	for i := range e.Storage {
		if e.Storage[i].UserID == id {
			tempSlice = append(tempSlice, e.Storage[i])
		}
	}

	if len(tempSlice) > 0 {
		return tempSlice
	}

	return nil
}
