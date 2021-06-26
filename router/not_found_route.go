package router

import (
	"net/http"

	"github.com/danangkonang/crud-rest/helper"
	"github.com/gorilla/mux"
)

func NotFoundRouter(r *mux.Router) {
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.MakeRespon(w, 404, "page not found", nil)
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.MakeRespon(w, http.StatusMethodNotAllowed, "Method NotAllowed", nil)
	})
}
