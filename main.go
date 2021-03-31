package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/danangkonang/crud-rest/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const STATIC_DIR = "/files/"

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix(STATIC_DIR).Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	PORT := os.Getenv("PORT")
	router.CrudRouter(r)

	fmt.Println("local server started at http://localhost:" + PORT)
	header := []string{
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Content-Type",
		"Authorization",
	}
	method := []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}
	origin := []string{"*"}
	http.ListenAndServe(":"+PORT, handlers.CORS(
		handlers.AllowedHeaders(header),
		handlers.AllowedMethods(method),
		handlers.AllowedOrigins(origin),
	)(r))
}
