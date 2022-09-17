package main

import (
	"github.com/tomoyasuzuki/sample-api/app/handler"
	"github.com/tomoyasuzuki/sample-api/app/repository"
	"github.com/tomoyasuzuki/sample-api/app/service"
	"log"
	"net/http"
)

const Port = ":80"

func main() {
	err := InitDB()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(Db)
	taskRepo := repository.NewTaskRepository(Db)
	userService := service.NewUserService(&userRepo)
	taskService := service.NewTaskService(&taskRepo)
	userHandler := handler.NewUserHandler(&userService)
	taskHandler := handler.NewTaskHandler(&taskService)

	srv := http.Server{
		Addr:    Port,
		Handler: Routes(userHandler, taskHandler),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
