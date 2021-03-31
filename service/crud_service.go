package service

import (
	"errors"
	"time"

	"github.com/danangkonang/crud-rest/config"
)

type Animal struct {
	ID          string    `json:"animal_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Color       string    `json:"color,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	CreateAt    time.Time `json:"created_at,omitempty"`
	UpdateAt    time.Time `json:"updated_at,omitempty"`
}

func SaveAnimal(animal *Animal) error {
	db := config.Connect()
	defer db.Close()
	query := `
		INSERT INTO
			animals (animal_id, name, color, description, image, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING animal_id
	`
	row := db.QueryRow(query, animal.ID, animal.Name, animal.Color, animal.Description, animal.Image, animal.CreateAt, animal.UpdateAt)
	err := row.Scan(&animal.ID)
	return err
}

func FindAnimal() ([]Animal, error) {
	db := config.Connect()
	defer db.Close()
	query := `
		SELECT
			*
		FROM
			animals
	`
	row, err := db.Query(query)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	var animal []Animal
	for row.Next() {
		var an Animal
		err := row.Scan(&an.ID, &an.Name, &an.Color, &an.Description, &an.Image, &an.CreateAt, &an.UpdateAt)
		if err != nil {
			return nil, err
		}
		animal = append(animal, an)
	}
	return animal, nil
}

func DetailAnimal(animal *Animal) error {
	db := config.Connect()
	defer db.Close()
	query := `
		SELECT
			*
		FROM
			animals
		WHERE animal_id = $1
	`
	row := db.QueryRow(query, animal.ID)
	err := row.Scan(&animal.ID, &animal.Name, &animal.Color, &animal.Description, &animal.Image, &animal.CreateAt, &animal.UpdateAt)
	return err
}

func DeleteAnimal(animal *Animal) error {
	db := config.Connect()
	defer db.Close()
	query := `
		DELETE FROM
			animals
		WHERE animal_id = $1
	`
	row := db.QueryRow(query, animal.ID)
	return row.Err()
}

func UpdateAnimal(animal *Animal) error {
	db := config.Connect()
	defer db.Close()
	query := `
		UPDATE
			animals SET name = $1, color = $2, description = $3, updated_at = $4
		WHERE animal_id = $5
	`
	row := db.QueryRow(query, animal.Name, animal.Color, animal.Description, animal.UpdateAt, animal.ID)
	return row.Err()
}
