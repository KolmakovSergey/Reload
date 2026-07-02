package models

import (
	"encoding/json"
	"time"
)

type ErrorDTO struct {
	Text      string    `json:"text"`
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
}

func NewErrorDTO(text string, code int) ErrorDTO {

	return ErrorDTO{
		Text:      text,
		Code:      code,
		Timestamp: time.Now(),
	}
}

func (e *ErrorDTO) MarshalErrorDTO() ([]byte, error) {
	errBytesDTO, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	return errBytesDTO, nil
}
