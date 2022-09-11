package main

import (
	"log"
	"net/http"
)

const Port = ":80"

func main() {
	err := InitDB()
	if err != nil {
		panic(err)
	}

	srv := http.Server{
		Addr:    Port,
		Handler: Routes(),
	}

	log.Fatal(srv.ListenAndServe())
}
