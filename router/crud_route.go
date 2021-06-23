package router

import (
	"github.com/danangkonang/crud-rest/config"
	"github.com/danangkonang/crud-rest/controller"
	"github.com/gorilla/mux"
)

func CrudRouter(router *mux.Router) {
	c := controller.NewAnimalHandler(config.NewDb())
	v1 := router.PathPrefix("/v1").Subrouter()

	v1.HandleFunc("/create", c.AnimalCreate).Methods("POST")
	// v1.HandleFunc("/show", controller.AnimalShow).Methods("GET")
	// v1.HandleFunc("/detail", controller.AnimalDetail).Methods("GET")
	// v1.HandleFunc("/edit", controller.AnimalEdit).Methods("PUT")
	// v1.HandleFunc("/delete", controller.AnimalDelete).Methods("DELETE")
}
