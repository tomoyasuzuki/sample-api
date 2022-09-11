package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func HandleGetTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")
	task, err := GetTask(taskID)
	if err != nil {
		fmt.Println(err)
	}
	RenderJSON(w, task)
}

func HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	var err error

	tag := r.FormValue("tag")
	user := r.FormValue("user")

	fmt.Println(tag)
	fmt.Println(user)

	if user != "" && tag != "" {
		tasks, err = GetTasksWithUserAndTag(user, tag)
	} else if tag != "" {
		tasks, err = GetTasksWithTag(tag)
	} else if user != "" {
		tasks, err = GetTasksWithUser(user)
	} else {
		tasks, err = GetTasks()
	}

	if err != nil {
		fmt.Println(err)
	}

	RenderJSON(w, tasks)
}

func HandlePostTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var err error
	if err = json.NewDecoder(r.Body).Decode(&task); err != nil {
		fmt.Println(err)
	}

	if err = PostTask(task); err != nil {
		fmt.Println(err)
	}

	RenderJSON(w, "succeed to post task")
}

func HandlePutTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	var err error
	if err = json.NewDecoder(r.Body).Decode(&task); err != nil {
		fmt.Println(err)
	}

	if err = UpdateTask(task); err != nil {
		fmt.Println(err)
	}

	RenderJSON(w, "succeed to update user")
}

func HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := GetUsers()
	if err != nil {
		fmt.Println(err)
	}
	RenderJSON(w, users)
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	user, err := GetUser(userID)
	if err != nil {
		fmt.Println(err)
	}
	RenderJSON(w, user)
}

func HandlePostUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var err error
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("Decode error")
		fmt.Println(err)
	}

	if err = PostUser(user); err != nil {
		fmt.Println(err)
	}

	RenderJSON(w, "succeed to post user")
}

func HandlePutUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var err error
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("Decode error")
		fmt.Println(err)
	}

	if err = UpdateUser(user); err != nil {
		fmt.Println(err)
	}

	RenderJSON(w, "succeed to post user")
}

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Route("/tasks", func(r chi.Router) {
		r.Get("/", HandleGetTasks)
		r.Get("/{id}", HandleGetTask)
		r.Post("/", HandlePostTask)
		r.Put("/", HandlePutTask)
	})

	mux.Route("/users", func(r chi.Router) {
		r.Get("/", HandleGetUsers)
		r.Get("/{id}", HandleGetUser)
		r.Post("/", HandlePostUser)
		r.Put("/", HandlePutUser)
	})

	return mux
}

// [{"id":3,"created_at":"2022-09-10T16:59:02Z","updated_at":"2022-09-10T16:59:02Z","deleted_at":null,"title":"Task 3","description":"This is a third task.","status":0,"assignees":[{"id":1,"created_at":"2022-09-10T16:58:49Z","updated_at":"2022-09-10T16:58:49Z","deleted_at":null,"name":"oshibori","tasks":null}],"tags":[{"id":1,"created_at":"2022-09-10T17:02:13Z","updated_at":"2022-09-10T17:02:13Z","deleted_at":null,"name":"test-tag-1","tasks":null}]}]
