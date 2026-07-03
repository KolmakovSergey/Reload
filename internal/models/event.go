package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	EventID    string
	UserID     int
	Activity   string
	ProductID  int
	HappenedAt time.Time
}

func NewEvent(userId int, activity string, productId int, happenedAt time.Time) (Event, error) {

	switch activity {
	case "view", "addToCart", "purchase":
		return Event{
			EventID:    uuid.New().String(),
			UserID:     userId,
			Activity:   activity,
			ProductID:  productId,
			HappenedAt: happenedAt,
		}, nil
	default:
		return Event{}, errors.New("undefined event")
	}

}

type EventDTO struct {
	UserID     int       `json:"userId"`
	Activity   string    `json:"action"`
	ProductID  int       `json:"productId"`
	HappenedAt time.Time `json:"timestamp"`
}

func NewEventDTO(userId int, activity string, productId int, happenedAt time.Time) EventDTO {

	return EventDTO{
		UserID:     userId,
		Activity:   activity,
		ProductID:  productId,
		HappenedAt: happenedAt,
	}
}
