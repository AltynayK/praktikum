package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type MyUrl struct {
	LongUrl  string `json:"LongUrl,omitempty"`
	ShortUrl int    `json:"ShortUrl,omitempty"`
}

var n1qlParams []interface{}

func PostEndpoint(w http.ResponseWriter, req *http.Request) {
	var url MyUrl
	_ = json.NewDecoder(req.Body).Decode(&url)

	n1qlParams = append(n1qlParams, url.LongUrl)

	url.ShortUrl = len(n1qlParams) - 1
	err := json.NewEncoder(w).Encode(url.ShortUrl)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}

}

func RootEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", PostEndpoint).Methods("POST")
	router.HandleFunc("/{id}", RootEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
