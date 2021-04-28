package repository

import (
	"database/sql"
	"fmt"
	"github.com/bogdanserdinov/tic-tac-toe-web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type AuthRepository struct{
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository{
	return &AuthRepository{
		db : db,
	}
}

func(r *AuthRepository) CreateUser(user tictactoe_web.User) (int,error){
	query := fmt.Sprintf("Insert into %s (name,password) values ($1,$2)",UserTable)

	result, err := r.db.Exec(query, user.Name,user.Password)
	if err != nil{
		return 0,nil
	}

	id, _ :=  result.LastInsertId()

	return int(id),nil
}

func(r *AuthRepository) GetUser(name string, password []byte) (tictactoe_web.User,error){
	var u tictactoe_web.User
	query := fmt.Sprintf("Select * from %s where name=$1 and password=$2",UserTable)

	row := r.db.QueryRow(query,name,password)

	err := row.Scan(&u.ID, &u.Name, &u.Password, &u.Role)
	if err != nil{
		logrus.Errorf("could not get data from db: %s",err.Error())
		return tictactoe_web.User{},nil
	}
	return u,nil
}