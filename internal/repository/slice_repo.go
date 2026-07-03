package repository

import (
	"errors"
	"reload/internal/models"
	"sync"
)

type EventRepo struct {
	Storage []models.Event
	mu      sync.RWMutex
}

func NewEventSliceRepo(events []models.Event) *EventRepo {
	return &EventRepo{
		Storage: events,
	}
}

func (e *EventRepo) SaveEvent(event models.EventDTO) error {

	newEvent, err := models.NewEvent(event.UserID, event.Activity, event.ProductID, event.HappenedAt)
	if err != nil {
		return errors.New("unable to save event with incorrect data")
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	e.Storage = append(e.Storage, newEvent)

	return nil
}

func (e *EventRepo) GetEventsByUserId(id int) ([]models.Event, error) {

	var tempSlice []models.Event
	e.mu.RLock()
	for i := range e.Storage {
		if e.Storage[i].UserID == id {
			tempSlice = append(tempSlice, e.Storage[i])
		}
	}
	e.mu.RUnlock()

	if len(tempSlice) > 0 {
		return tempSlice, nil
	}

	return []models.Event{}, nil
}

func (e *EventRepo) GetAllEvents() ([]models.Event, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	eventsCopy := make([]models.Event, len(e.Storage))
	copy(eventsCopy, e.Storage)

	return eventsCopy, nil
}
