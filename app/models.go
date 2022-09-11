package main

import (
	"time"
)

// NOTE: IDはUUID型の方が望ましい気がするけど、開発時に楽なので一旦uint型で対処する

type Task struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      uint       `json:"status"`
	Assignees   []*User    `json:"assignees" gorm:"many2many:task_users;"`
	Tags        []*Tag     `json:"tags" gorm:"many2many:task_tags;"`
}

type Tag struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
	Tasks     []*Task    `json:"tasks" gorm:"many2many:task_tags;"`
}

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name"`
	Tasks     []*Task    `json:"tasks" gorm:"many2many:task_users;"`
}
