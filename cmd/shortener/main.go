package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type MyUrl struct {
	ShortUrl int
	LongUrl  string
}

func PostEndpoint(w http.ResponseWriter, req *http.Request) {
	var url MyUrl
	_ = json.NewDecoder(req.Body).Decode(&url)
	var aParams []interface{}
	aParams = append(aParams, url.LongUrl)
	for i := 0; i < len(aParams); i++ {
		if aParams[i] == url.LongUrl {
			url.ShortUrl = i
		}
	}
	json.NewEncoder(w).Encode(url.ShortUrl)
}

func GetEndpoint(w http.ResponseWriter, req *http.Request) {
	var aParams []interface{}
	params := req.URL.Query()
	aParams = append(aParams, params.Get("shortUrl"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", PostEndpoint)
	mux.HandleFunc("/{id}", GetEndpoint)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
