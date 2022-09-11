package main

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func VerifyToken() {

}

func GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(hash), err
	}
	return string(hash), nil
}

func CompareHash(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func SignUP(userName string, email string, password string) error {
	var err error
	user := User{
		Name:      userName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		Tasks:     []*Task{},
	}

	if err = Db.Create(&user).Error; err != nil {
		return err
	}

	hash, err := GenerateHash(password)
	if err != nil {
		return err
	}

	cred := Credentials{
		UserID:   user.ID,
		Email:    email,
		Password: hash,
	}

	if err = Db.Create(&cred).Error; err != nil {
		return err
	}

	return nil
}

func SignIn(w http.ResponseWriter, r *http.Request) {

}
