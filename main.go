package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/danangkonang/crud-rest/config"
	"github.com/danangkonang/crud-rest/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/files/").Handler(
		http.StripPrefix(
			"/files/",
			http.FileServer(
				http.Dir("."+"/files/"),
			),
		),
	)
	router.CrudRouter(r, config.NewDb())
	router.NotFoundRouter(r)

	header := []string{
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Content-Type",
		"Authorization",
	}
	method := []string{"GET", "POST", "PUT", "DELETE"}
	origin := []string{"*"}

	srv := &http.Server{
		Addr:         "127.0.0.1:9000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: handlers.CORS(
			handlers.AllowedHeaders(header),
			handlers.AllowedMethods(method),
			handlers.AllowedOrigins(origin),
		)(r),
	}
	fmt.Println("local server started at http://" + srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
