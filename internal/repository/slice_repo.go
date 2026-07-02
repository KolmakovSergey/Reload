package repository

import (
	"reload/internal/models"
	"sync"
)

type EventRepo struct {
	Storage []models.Event
	mut     sync.RWMutex
}

func NewEventSliceRepo(events []models.Event) *EventRepo {
	return &EventRepo{
		Storage: events,
	}
}

func (e *EventRepo) SaveEvent(event models.Event) error {
	e.mut.Lock()
	defer e.mut.Unlock()
	e.Storage = append(e.Storage, event)
	
	return nil
}

func (e *EventRepo) GetEventsByUserId(id int) ([]models.Event, error) {

	var tempSlice []models.Event
	e.mut.RLock()
	for i := range e.Storage {
		if e.Storage[i].UserID == id {
			tempSlice = append(tempSlice, e.Storage[i])
		}
	}
	e.mut.RUnlock()

	if len(tempSlice) > 0 {
		return tempSlice, nil
	}

	return []models.Event{}, nil
}

func (e *EventRepo) GetAllEvents() ([]models.Event, error) {
	e.mut.RLock()
	defer e.mut.RUnlock()

	eventsCopy := make([]models.Event, len(e.Storage))
    copy(eventsCopy, e.Storage)
	
	return eventsCopy, nil
}
