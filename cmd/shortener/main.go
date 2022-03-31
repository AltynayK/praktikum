package main

import (
	"log"
	"net/http"

	"github.com/AltynayK/praktikum/internal/handler"
	"github.com/AltynayK/praktikum/internal/service"
	"github.com/gorilla/mux"
)

const port = ":8080"

func main() {
	mux := initHandlers()
	service.IdList = make(map[int]string)

	srv := http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("App listening port: %s", port)
	log.Fatal(srv.ListenAndServe())
}

func initHandlers() *mux.Router {
	// TODO: how handler 404 (if not found some url, example: /not_exist_url)
	// TODO: handle "Not Allowed Method" example: DELETE method request to /

	router := mux.NewRouter()
	router.HandleFunc("/", handler.Post).Methods("POST")
	router.HandleFunc("/{id}", handler.Get).Methods("GET")

	return router
}
