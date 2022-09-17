package repository

import (
	appError "github.com/tomoyasuzuki/sample-api/app/error"
	"github.com/tomoyasuzuki/sample-api/app/model"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetTask(id string) (model.Task, error)
	GetTasksWith(userName string, tagName string) ([]model.Task, error)
	CreateTask(task model.Task) error
	UpdateTask(task model.Task) error
}

type TaskRepository struct {
	Db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{db}
}

func (t *TaskRepository) GetTask(id string) (model.Task, error) {
	var task model.Task

	if err := t.Db.
		Model(&model.Task{}).
		Preload("Assignees").
		Preload("Tags").
		First(&task, id).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (t *TaskRepository) GetTasks() ([]model.Task, error) {
	var tasks []model.Task

	if err := t.Db.
		Model(&model.Task{}).
		Preload("Assignees").
		Preload("Tags").
		Find(&tasks).Error; err != nil {
		return tasks, appError.New(appError.Unknown)
	}
	return tasks, nil
}

func (t *TaskRepository) GetTasksWith(userName string, tagName string) ([]model.Task, error) {
	var tasks []model.Task

	if userName != "" && tagName != "" {
		if err := t.getTaskWith(
			tasks,
			withUser(t.Db),
			withTag(t.Db),
			byUserName(t.Db, userName),
			byTagName(t.Db, tagName)); err != nil {
			return tasks, err
		}
	} else if userName != "" {
		if err := t.getTaskWith(
			tasks,
			withUser(t.Db),
			byUserName(t.Db, userName)); err != nil {
			return tasks, err
		}
	} else if tagName != "" {
		if err := t.getTaskWith(
			tasks,
			withTag(t.Db),
			byTagName(t.Db, tagName)); err != nil {
			return tasks, err
		}
	} else {
		if err := t.getTaskWith(tasks); err != nil {
			return tasks, err
		}
	}

	return tasks, nil
}

func (t *TaskRepository) CreateTask(task model.Task) error {
	if err := t.Db.Create(&task).Error; err != nil {
		return appError.New(appError.Unknown)
	}

	return nil
}

func (t *TaskRepository) UpdateTask(task model.Task) error {
	if err := t.Db.Save(&task).Error; err != nil {
		return appError.New(appError.Unknown)
	}

	return nil
}

func (t *TaskRepository) getTaskWith(tasks []model.Task, scopes ...func(*gorm.DB) *gorm.DB) error {
	if err := t.Db.
		Model(&model.Task{}).
		Preload("Assignees").
		Preload("Tags").
		Scopes(scopes...).
		Find(&tasks).Error; err != nil {
		return appError.New(appError.Unknown)
	}
	return nil
}

func withUser(db *gorm.DB) func(*gorm.DB) *gorm.DB {
	return func(*gorm.DB) *gorm.DB {
		return db.
			Joins("JOIN task_users ON tasks.id = task_users.task_id").
			Joins("JOIN users ON users.id = task_users.user_id")
	}
}

func withTag(db *gorm.DB) func(*gorm.DB) *gorm.DB {
	return func(*gorm.DB) *gorm.DB {
		return db.
			Joins("JOIN task_tags ON tasks.id = task_tags.task_id").
			Joins("JOIN tags ON tags.id = task_tags.tag_id")
	}
}

func byUserName(db *gorm.DB, userName string) func(*gorm.DB) *gorm.DB {
	return func(*gorm.DB) *gorm.DB {
		return db.Where("users.name = ?", userName)
	}
}

func byTagName(db *gorm.DB, tagName string) func(*gorm.DB) *gorm.DB {
	return func(*gorm.DB) *gorm.DB {
		return db.Where("tags.name = ?", tagName)
	}
}
