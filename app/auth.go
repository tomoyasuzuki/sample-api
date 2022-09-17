package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tomoyasuzuki/sample-api/app/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func VerifyToken() {

}

type SignUpParam struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

func SignUp(param SignUpParam) (string, error) {
	var err error
	user := model.User{
		Name:      param.UserName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
		Token:     nil,
		Tasks:     []*model.Task{},
	}

	if err = Db.Create(&user).Error; err != nil {
		return "", err
	}

	hash, err := GenerateHash(param.Password)
	if err != nil {
		return "", err
	}

	cred := model.Credentials{
		UserID:   user.ID,
		Email:    param.Email,
		Password: hash,
	}

	if err = Db.Create(&cred).Error; err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"nbf":    time.Now().Unix(),
		"iat":    time.Now().Unix(),
		"exp":    time.Now().Add(time.Hour * 24 * 30),
	})

	ss, err := token.SignedString([]byte("SECRET_KEY"))

	if err != nil {
		return "", err
	}

	return ss, nil
}

func SignIn(param SignInParam) (model.User, error) {
	var err error
	var cred model.Credentials
	var user model.User

	if err != nil {
		fmt.Println(err)
		return user, err
	}

	if err = Db.Where("email = ?", param.Email).First(&cred).Error; err != nil {
		fmt.Println(err)
		return user, err
	}

	hash, err := GenerateHash(param.Password)

	if err = CompareHash(hash, param.Password); err != nil {
		fmt.Println(err)
		return user, err
	}

	if err = Db.First(&user, cred.UserID).Error; err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}
