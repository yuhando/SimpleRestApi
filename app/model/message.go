package model

import (
	"errors"
)

// Message Model
type Message struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

// Validate model Message
func (m *Message) Validate() error {
	if m.Name == "" {
		return errors.New("Required Name")
	}
	if m.Message == "" {
		return errors.New("Required Message")
	}

	return nil
}
