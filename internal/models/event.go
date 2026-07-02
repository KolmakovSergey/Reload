package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	EventID   string
	UserID    int
	Action    string
	ProductID int
	Timestamp time.Time
}

type EventDTO struct {
	UserID    int       `json:"userId"`
	Action    string    `json:"action"`
	ProductID int       `json:"productId"`
	Timestamp time.Time `json:"timestamp"`
}

func NewEvent(userId int, action string, productId int, timestamp time.Time) (Event, error) {

	switch action {
	case "view", "addToCart", "purchase":
		return Event{
			EventID:   uuid.New().String(),
			UserID:    userId,
			Action:    action,
			ProductID: productId,
			Timestamp: timestamp,
		}, nil
	default:
		return Event{}, errors.New("undefined event")
	}

}
