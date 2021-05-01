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

	rows,err := r.db.Query(query,name,password)
	if err != nil {
		logrus.Errorf("unsuccesful query to db: %s",err.Error())
	}

	for rows.Next(){
		err = rows.Scan(&u.ID, &u.Name, &u.Password, &u.Role)
		fmt.Println(u)
		if err != nil{
			logrus.Errorf("could not get data from db: %s",err.Error())
			return tictactoe_web.User{},err
		}
	}

	return u,nil
}