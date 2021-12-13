package model

import "time"

type Animal struct {
	ID          int       `json:"animal_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Color       string    `json:"color,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type ResponseAnimal struct {
	ID          int       `json:"animal_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Color       string    `json:"color,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
