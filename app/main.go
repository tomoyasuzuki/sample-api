package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World")
}

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handleRoot)

	return mux
}

func main() {
	srv := http.Server{
		Addr:    ":80",
		Handler: routes(),
	}

	log.Fatal(srv.ListenAndServe())
}
