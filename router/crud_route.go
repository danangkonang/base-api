package router

import (
	"github.com/danangkonang/crud-rest/controller"
	"github.com/gorilla/mux"
)

func CrudRouter(router *mux.Router) {
	prefix := "/v1/animal"
	router.HandleFunc(prefix+"/create", controller.AnimalCreate).Methods("POST")
	router.HandleFunc(prefix+"/show", controller.AnimalShow).Methods("GET")
	router.HandleFunc(prefix+"/detail", controller.AnimalDetail).Methods("GET")
	router.HandleFunc(prefix+"/edit", controller.AnimalEdit).Methods("PUT")
	router.HandleFunc(prefix+"/delete", controller.AnimalDelete).Methods("DELETE")
}
