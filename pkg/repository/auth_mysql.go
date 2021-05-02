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
	if r.db == nil{
		logrus.Fatalf("db = nil")
	}
	var id int64
	query := fmt.Sprintf("Insert into %s (name,password) values (?,?)",UserTable)

	result, err := r.db.Exec(query, user.Name,user.Password)

	if err != nil{
		logrus.Fatalf( "could not get id of inserted row: %s",err.Error())
		return 0,err
	}
	id,_ = result.LastInsertId()
	return int(id),nil
}

func(r *AuthRepository) GetUser(name string, password string) (tictactoe_web.User,error){
	var u tictactoe_web.User
	query := fmt.Sprintf("Select * from %s where name=? and password=?",UserTable)
	rows:= r.db.QueryRow(query,name,password)

	err := rows.Scan(&u.ID, &u.Name, &u.Password, &u.Role)
	if err != nil {
		logrus.Errorf("could not get row from db: %s",err.Error())
	}

	return u,nil
}