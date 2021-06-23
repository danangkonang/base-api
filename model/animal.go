package model

import "time"

type Animal struct {
	ID          string    `json:"animal_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Color       string    `json:"color,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty"`
	UpdateAt    time.Time `json:"updated_at,omitempty"`
}
