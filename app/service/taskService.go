package service

import (
	"github.com/tomoyasuzuki/sample-api/app/model"
	"github.com/tomoyasuzuki/sample-api/app/repository"
)

type ITaskService interface {
	GetTask(id string) (model.Task, error)
	GetTasks(userName string, tagName string) ([]model.Task, error)
	CreateTask(task model.Task) error
	UpdateTask(task model.Task) error
}

type TaskService struct {
	repo repository.ITaskRepository
}

func NewTaskService(repo *repository.ITaskRepository) ITaskService {
	return &TaskService{*repo}
}

func (ts *TaskService) GetTask(id string) (model.Task, error) {
	return ts.repo.GetTask(id)
}

func (ts *TaskService) GetTasks(userName string, tagName string) ([]model.Task, error) {
	return ts.repo.GetTasksWith(userName, tagName)
}

func (ts *TaskService) CreateTask(task model.Task) error {
	return ts.repo.CreateTask(task)
}

func (ts *TaskService) UpdateTask(task model.Task) error {
	return ts.repo.UpdateTask(task)
}
