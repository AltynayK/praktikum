package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/AltynayK/praktikum/internal/domain"
	"github.com/AltynayK/praktikum/internal/service"
	"github.com/gorilla/mux"
)

func Post(rw http.ResponseWriter, req *http.Request) {
	var (
		url     *domain.Url
		jsonRes []byte
		err     error
	)

	if err = json.NewDecoder(req.Body).Decode(&url); err != nil {
		log.Println("Post req Error:", err)
		errRes := domain.ErrResponse{
			Msg: fmt.Sprintf("json decode err: %s", err),
		}

		rw.WriteHeader(400)
		json.NewEncoder(rw).Encode(errRes)

		return
	}

	//log.Printf("parsed url: %v\n", url)

	// TODO: some service logic
	id := service.WriteURLByID(url.LongURL)

	okRes := domain.PostResponse{
		ID: id,
	}

	if jsonRes, err = json.Marshal(okRes); err != nil {
		rw.WriteHeader(500)
		fmt.Fprintf(rw, "response json marshal err")

		return
	}

	// set "Created" status 201
	rw.WriteHeader(201)
	fmt.Fprint(rw, string(jsonRes))
}

func Get(rw http.ResponseWriter, req *http.Request) {
	var jsonRes []byte

	vars := mux.Vars(req)
	id := vars["id"]

	ID, err := strconv.Atoi(id)
	if err != nil {
		errRes := domain.ErrResponse{
			Msg: err.Error(),
		}

		rw.WriteHeader(400)
		json.NewEncoder(rw).Encode(errRes)

		return
	}

	//log.Println("parsed ID:", id)
	url := service.GetURLFromID(ID)

	okRes := domain.GetResponse{
		Url: url,
	}

	if jsonRes, err = json.Marshal(okRes); err != nil {
		rw.WriteHeader(500)
		fmt.Fprintf(rw, "response json marshal err")

		return
	}

	// set "Created" status 201
	rw.WriteHeader(307)
	fmt.Fprint(rw, string(jsonRes))
}
