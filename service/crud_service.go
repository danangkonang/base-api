package service

import (
	"database/sql"
	"errors"

	"github.com/danangkonang/crud-rest/config"
	"github.com/danangkonang/crud-rest/model"
)

type AnimalService interface {
	SaveAnimal(animal *model.Animal) error
	FindAnimal() ([]model.Animal, error)
	DetailAnimal(id string) (model.Animal, error)
	DeleteAnimal(animal *model.Animal) error
	UpdateAnimal(animal *model.Animal) error
}

func NewServiceAnimal(Con *config.DB) AnimalService {
	return &PsqlAnimalService{
		Psql: Con.Postgresql,
	}
}

type PsqlAnimalService struct {
	Psql *sql.DB
}

func (m *PsqlAnimalService) SaveAnimal(animal *model.Animal) error {
	// defer m.Psql.Close()
	query := `
		INSERT INTO
			animals (animal_id, name, color, description, image, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING animal_id
	`
	row := m.Psql.QueryRow(query, animal.ID, animal.Name, animal.Color, animal.Description, animal.Image, animal.CreateAt, animal.UpdateAt)
	err := row.Scan(&animal.ID)
	return err
}

func (m *PsqlAnimalService) FindAnimal() ([]model.Animal, error) {
	// defer m.Psql.Close()
	query := `
		SELECT
			*
		FROM
			animals
	`
	row, err := m.Psql.Query(query)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	var animal []model.Animal
	for row.Next() {
		var an model.Animal
		err := row.Scan(&an.ID, &an.Name, &an.Color, &an.Description, &an.Image, &an.CreateAt, &an.UpdateAt)
		if err != nil {
			return nil, err
		}
		animal = append(animal, an)
	}
	defer row.Close()
	return animal, nil
}

func (m *PsqlAnimalService) DetailAnimal(uid string) (model.Animal, error) {
	// defer m.Psql.Close()
	var animal model.Animal
	query := `
		SELECT
			animal_id, name, color, description, image, created_at, updated_at
		FROM
			animals
		WHERE animal_id = $1
	`
	row := m.Psql.QueryRow(query, uid)
	// fmt.Println(row)
	err := row.Scan(&animal.ID, &animal.Name, &animal.Color, &animal.Description, &animal.Image, &animal.CreateAt, &animal.UpdateAt)
	return animal, err
}

func (m *PsqlAnimalService) DeleteAnimal(animal *model.Animal) error {
	// defer m.Psql.Close()
	query := `
		DELETE FROM
			animals
		WHERE animal_id = $1
	`
	row := m.Psql.QueryRow(query, animal.ID)
	return row.Err()
}

func (m *PsqlAnimalService) UpdateAnimal(animal *model.Animal) error {
	// defer m.Psql.Close()
	query := `
		UPDATE
			animals SET name = $1, color = $2, description = $3, updated_at = $4
		WHERE animal_id = $5
	`
	row := m.Psql.QueryRow(query, animal.Name, animal.Color, animal.Description, animal.UpdateAt, animal.ID)
	return row.Err()
}
