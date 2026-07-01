package models

import (
	"errors"
	"math/rand"
	"time"
)

type Event struct {
	EventID   int
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
			EventID:   rand.Int(),
			UserID:    userId,
			Action:    action,
			ProductID: productId,
			Timestamp: timestamp,
		}, nil
	default:
		return Event{}, errors.New("undefined event")
	}

}
