package router

import (
	"github.com/danangkonang/crud-rest/config"
	"github.com/danangkonang/crud-rest/controller"
	"github.com/gorilla/mux"
)

func CrudRouter(router *mux.Router, db *config.DB) {
	c := controller.NewAnimalHandler(db)
	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/animals", c.AnimalShow).Methods("GET")
	v1.HandleFunc("/animal", c.AnimalCreate).Methods("POST")
	v1.HandleFunc("/animal", c.AnimalDetail).Methods("GET")
	v1.HandleFunc("/animal", c.AnimalEdit).Methods("PUT")
	v1.HandleFunc("/animal", c.AnimalDelete).Methods("DELETE")
}
