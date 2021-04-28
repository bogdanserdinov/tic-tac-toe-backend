package service

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService{
	return &AuthService{
		repo : repo,
	}
}

func(a *AuthService) CreateUser(user tictactoe_web.User) (int,error){
	user.Password = a.HashPassword(user.Password)
	return a.repo.CreateUser(user)
}


func (a *AuthService) HashPassword(password string) string{
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil{
		logrus.Errorf("could not bcrypt password: %s",err.Error())
	}
	return string(hashedPassword)
}