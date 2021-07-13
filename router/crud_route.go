package router

import (
	"net"
	"net/http"

	"github.com/danangkonang/crud-rest/config"
	"github.com/danangkonang/crud-rest/controller"
	"github.com/danangkonang/crud-rest/helper"
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

	v1.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			helper.MakeRespon(w, 400, "", err.Error())
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					helper.MakeRespon(w, 200, "", ipnet.IP.String())
				}
			}
		}
	})

}
