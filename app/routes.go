package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/tomoyasuzuki/sample-api/app/handler"
	"net/http"
)

func Routes(userHandler handler.IUserHandler, taskHandler handler.ITaskHandler) http.Handler {
	mux := chi.NewRouter()

	// mux.Post("/signup", )
	// mux.Post("/login",)

	mux.Route("/tasks", func(r chi.Router) {
		r.Get("/", taskHandler.GetTasks)
		r.Get("/{id}", taskHandler.GetTask)
		r.Post("/", taskHandler.PostTask)
		r.Put("/", taskHandler.PutTask)
	})

	mux.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Get("/{id}", userHandler.GetUser)
		r.Post("/", userHandler.PostUser)
		r.Put("/", userHandler.PutUser)
	})

	return mux
}
