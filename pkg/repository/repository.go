package repository

import "database/sql"

type Authorization interface {
	SignIn()
	SignUp()
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository{
	return &Repository{}
}
