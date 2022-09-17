package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	appError "github.com/tomoyasuzuki/sample-api/app/error"
	"github.com/tomoyasuzuki/sample-api/app/model"
	"github.com/tomoyasuzuki/sample-api/app/service"
	"net/http"
)

type IUserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	PostUser(w http.ResponseWriter, r *http.Request)
	PutUser(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler(service *service.IUserService) IUserHandler {
	return &UserHandler{*service}
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	user, err := uh.service.GetUser(userID)
	if err != nil {

	}

	if err == nil {
		Response(w, user, nil)
	} else {
		appErr := err.(appError.AppError)
		Response(w, user, &appErr)
	}
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uh.service.GetUsers()
	if err != nil {
		fmt.Println(err)
	}

	if err == nil {
		Response(w, users, nil)
	} else {
		appErr := err.(appError.AppError)
		Response(w, users, &appErr)
	}
}

func (uh *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var err error
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {

	}

	if err = uh.service.CreateUser(user); err != nil {
		fmt.Println(err)
	}

	if err == nil {
		Response(w, nil, nil)
	} else {
		appErr := err.(appError.AppError)
		Response(w, nil, &appErr)
	}
}

func (uh *UserHandler) PutUser(w http.ResponseWriter, r *http.Request) {}
