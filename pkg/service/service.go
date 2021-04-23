package service

import "github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"

type Authorization interface {
	SignIn()
	SignUp()
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service{
	return &Service{}
}