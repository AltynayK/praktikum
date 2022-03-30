package main

import (
	"go/format"
	"log"

	"github.com/AltynayK/praktikum"
)

func main() {
	srv := new(praktikum.Server)
	if err:=srv.Run(port: "8080"); err !=nil {
		log.Fatal(format: "error occured while running http server: %s", err.Error())
	}
}
