package main

import (
	"fmt"
	appError "github.com/tomoyasuzuki/sample-api/app/error"
	"github.com/tomoyasuzuki/sample-api/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

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
		return appError.New(appError.ConnectionFailed)
	}

	fmt.Println("Succeeded to connect database.")

	if err = Db.AutoMigrate(&model.Task{}); err != nil {
		return appError.New(appError.ConnectionFailed)
	}
	if err = Db.AutoMigrate(&model.User{}); err != nil {
		return appError.New(appError.ConnectionFailed)
	}
	if err = Db.AutoMigrate(&model.Tag{}); err != nil {
		return appError.New(appError.ConnectionFailed)
	}
	if err = Db.AutoMigrate(&model.Credentials{}); err != nil {
		return appError.New(appError.ConnectionFailed)
	}

	return nil
}
