package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func createDummyData() {
	user := User{
		Name:  "oshibori",
		Tasks: []*Task{},
	}
	task := Task{
		Title:       "test task",
		Description: "This is a test task.",
		Assignees:   []*User{},
		Tags:        []*Tag{},
	}

	Db.Create(&user)
	Db.Create(&task)
}

func InitDB() error {
	var err error
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")
	hostName := os.Getenv("MYSQL_HOSTNAME")
	port := os.Getenv("MYSQL_PORT")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		user,
		pass,
		hostName,
		port,
		dbName)
	retryCount := 1
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	for err != nil && retryCount < 4 {
		fmt.Printf("%d:failed to connect database.", retryCount)
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		retryCount += 1
	}

	if err != nil {
		return err
	}

	fmt.Println("Succeeded to connect database.")

	if err = Db.AutoMigrate(&Task{}); err != nil {
		return err
	}
	if err = Db.AutoMigrate(&User{}); err != nil {
		return err
	}
	if err = Db.AutoMigrate(&Tag{}); err != nil {
		return err
	}

	return nil
}
