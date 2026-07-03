package server

import (
	"time"
)

type ErrorDTO struct {
	Text       string    `json:"text"`
	Code       int       `json:"code"`
	HappenedAt time.Time `json:"happenedAt"`
}

func NewErrorDTO(text string, code int) ErrorDTO {

	return ErrorDTO{
		Text:      text,
		Code:      code,
		HappenedAt: time.Now(),
	}
}
