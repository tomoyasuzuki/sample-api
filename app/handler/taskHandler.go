package handler

import (
	"github.com/go-chi/chi/v5"
	appError "github.com/tomoyasuzuki/sample-api/app/error"
	"github.com/tomoyasuzuki/sample-api/app/service"
	"net/http"
)

type ITaskHandler interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	GetTasks(w http.ResponseWriter, r *http.Request)
	PostTask(w http.ResponseWriter, r *http.Request)
	PutTask(w http.ResponseWriter, r *http.Request)
}

type TaskHandler struct {
	service service.ITaskService
}

func NewTaskHandler(service *service.ITaskService) ITaskHandler {
	return &TaskHandler{*service}
}

func (th *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	taskID := chi.URLParam(r, "id")
	task, err := th.service.GetTask(taskID)
	if err != nil {

	}

	if err == nil {
		Response(w, task, nil)
	} else {
		appErr := err.(appError.AppError)
		Response(w, task, &appErr)
	}
}

func (th *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	var user string
	var tag string

	user = r.FormValue("user")
	tag = r.FormValue("tag")

	tasks, err := th.service.GetTasks(user, tag)
	if err != nil {

	}

	if err == nil {
		Response(w, tasks, nil)
	} else {
		appErr := err.(appError.AppError)
		Response(w, tasks, &appErr)
	}
}

func (th *TaskHandler) PostTask(w http.ResponseWriter, r *http.Request) {}
func (th *TaskHandler) PutTask(w http.ResponseWriter, r *http.Request)  {}
