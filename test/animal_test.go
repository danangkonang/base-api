package test

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/danangkonang/crud-rest/config"
	"github.com/danangkonang/crud-rest/controller"
	"github.com/danangkonang/crud-rest/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' connection", err)
	}
	return sqlDB, mock
}

type DB struct {
	Postgresql *sql.DB
}

func Server(db *config.DB) *mux.Router {
	router := mux.NewRouter()
	c := controller.NewAnimalHandler(db)
	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/animals", c.AnimalShow).Methods("GET")
	v1.HandleFunc("/animal", c.AnimalCreate).Methods("POST")
	v1.HandleFunc("/animal", c.AnimalDetail).Methods("GET")
	v1.HandleFunc("/animal", c.AnimalEdit).Methods("PUT")
	v1.HandleFunc("/animal", c.AnimalDelete).Methods("DELETE")
	return router
}

func TestFindAnimals(t *testing.T) {
	sqlDB, mock := NewMock()
	con := &DB{
		Postgresql: sqlDB,
	}
	defer con.Postgresql.Close()
	timestamp := time.Now()
	rows := sqlmock.NewRows([]string{"animal_id", "name", "color", "description", "image", "created_at", "updated_at"}).AddRow(1, "gajah", "blue", "long nose", "gajah.jpg", timestamp, timestamp)
	query := `
		SELECT
			animal_id, name, color, description, image, created_at, updated_at
		FROM
			animals
	`
	mock.ExpectQuery(query).WillReturnRows(rows)
	request, _ := http.NewRequest("GET", "/v1/animals", nil)
	response := httptest.NewRecorder()
	Server((*config.DB)(con)).ServeHTTP(response, request)
	body, _ := ioutil.ReadAll(response.Body)
	expectedResponse := fmt.Sprintf("{\"status\":200,\"message\":\"success\",\"data\":[{\"animal_id\":1,\"name\":\"gajah\",\"color\":\"blue\",\"description\":\"long nose\",\"image\":\"gajah.jpg\",\"created_at\":\"%s\",\"updated_at\":\"%s\"}]}", timestamp.Format(time.RFC3339Nano), timestamp.Format(time.RFC3339Nano))
	assert.Equal(t, http.StatusOK, response.Code, "Invalid response code")
	assert.Equal(t, expectedResponse, string(bytes.TrimSpace(body)))
}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestCreateAnimal(t *testing.T) {
	sample := []struct {
		body    model.Animal
		status  int
		message string
	}{
		{
			body: model.Animal{
				Name:        "",
				Color:       "",
				Description: "",
				Image:       "",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			status:  200,
			message: "success",
		},
	}
	sqlDB, mock := NewMock()
	con := &DB{
		Postgresql: sqlDB,
	}
	defer con.Postgresql.Close()
	for _, smpl := range sample {
		sqlmock.NewRows([]string{"animal_id", "name", "color", "description", "image", "created_at", "updated_at"}).AddRow(1, smpl.body.Name, smpl.body.Color, smpl.body.Description, smpl.body.Image, smpl.body.CreatedAt, smpl.body.UpdatedAt)
		query := "INSERT INTO animals \\(name, color, description, image, created_at, updated_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6\\)"
		bd, _ := json.Marshal(smpl.body)
		mock.ExpectExec(query).WithArgs(smpl.body.Name, smpl.body.Color, smpl.body.Description, smpl.body.Image, AnyTime{}, AnyTime{}).WillReturnResult(sqlmock.NewResult(0, 1))
		request, _ := http.NewRequest("POST", "/v1/animal", bytes.NewReader(bd))
		request.Header.Add("Content-Type", "application/json")
		response := httptest.NewRecorder()

		Server(&config.DB{Postgresql: con.Postgresql}).ServeHTTP(response, request)

		responseBody := make(map[string]interface{})
		rbody, _ := ioutil.ReadAll(response.Body)
		err := json.Unmarshal(rbody, &responseBody)
		if err != nil {
			t.Errorf("can not conver to json: %v", err)
		}
		assert.Equal(t, smpl.status, response.Code, "")
		assert.Equal(t, smpl.message, responseBody["message"], "")
	}
}

func TestDelete(t *testing.T) {
	sqlDB, mock := NewMock()
	con := &DB{
		Postgresql: sqlDB,
	}
	defer con.Postgresql.Close()

	body := model.Animal{
		ID: 1,
	}
	bd, _ := json.Marshal(body)
	query := "DELETE FROM animals WHERE animal_id = \\$1"
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	request, _ := http.NewRequest("DELETE", "/v1/animal", bytes.NewReader(bd))
	request.Header.Add("Content-Type", "application/json")
	response := httptest.NewRecorder()

	Server(&config.DB{Postgresql: con.Postgresql}).ServeHTTP(response, request)

	responseBody := make(map[string]interface{})
	rbody, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(rbody, &responseBody)
	if err != nil {
		t.Errorf("can not conver to json: %v", err)
	}
	assert.Equal(t, 200, response.Code, "")
	assert.Equal(t, "success", responseBody["message"], "")
}

func TestUpdate(t *testing.T) {
	sqlDB, mock := NewMock()
	con := &DB{
		Postgresql: sqlDB,
	}
	defer con.Postgresql.Close()

	body := model.Animal{
		ID:          1,
		Name:        "",
		Color:       "",
		Description: "",
		Image:       "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	bd, _ := json.Marshal(body)
	query := "UPDATE animals SET name = \\$1, color = \\$2, description = \\$3, updated_at = \\$4 WHERE animal_id = \\$5"
	mock.ExpectExec(query).WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	request, _ := http.NewRequest("UPDATE", "/v1/animal", bytes.NewReader(bd))
	request.Header.Add("Content-Type", "application/json")
	response := httptest.NewRecorder()

	Server(&config.DB{Postgresql: con.Postgresql}).ServeHTTP(response, request)

	responseBody := make(map[string]interface{})
	rbody, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(rbody, &responseBody)
	if err != nil {
		t.Errorf("can not conver to json: %v", err)
	}
	assert.Equal(t, 200, response.Code, "")
	assert.Equal(t, "success", responseBody["message"], "")
}
