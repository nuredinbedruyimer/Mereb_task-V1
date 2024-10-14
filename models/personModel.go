package models

import "github.com/google/uuid"

type Person struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name" validate:"required"`
	Age     int64     `json:"age" validate:"required"`
	Hobbies []string  `json:"hobbies"`
}
