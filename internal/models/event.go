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
    Meta      map[string]interface{}
}

func NewEvent(userId int, action string, productId int, meta map[string]interface{}) (Event, error) {

	switch action {
	case "view", "add_to_cart", "purchase":
		return Event{
			EventID:   rand.Int(),
			UserID:    userId,
			Action:     action,
			ProductID: productId,
			Timestamp:  time.Now(),
			Meta:       meta,
		}, nil
	default:
		return Event{}, errors.New("undefined event")
	}

}
