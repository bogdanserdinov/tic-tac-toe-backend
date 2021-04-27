package service

import (
	tictactoe_web "github.com/bogdanserdinov/tic-tac-toe-web"
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

func(a *AuthService) SignIn(user tictactoe_web.User) (int,error){
	return 0,nil
}

func (a *AuthService) HashPassword(password []byte) []byte{
	hashedPassword,err := bcrypt.GenerateFromPassword(password,bcrypt.DefaultCost)
	if err != nil{
		logrus.Errorf("could not bcrypt password: %s",err.Error())
	}
	return hashedPassword
}