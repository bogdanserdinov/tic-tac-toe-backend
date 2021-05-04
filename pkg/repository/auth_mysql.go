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
		logrus.Fatalf( "could execute insert query: %s",err.Error())
		return 0,err
	}
	//create a record in stats table with id of user
	id,_ = result.LastInsertId()
	query = fmt.Sprintf("Insert into %s (id,total,wins,draws,losses) values (?,0,0,0,0)", UserStatsTable)
	_, err = r.db.Exec(query,int(id))
	if err != nil{
		logrus.Fatalf( "could execute insert query: %s",err.Error())
		return 0,err
	}

	return int(id),nil
}

func(r *AuthRepository) GetUser(name string, password string) (tictactoe_web.User,error){
	var user tictactoe_web.User
	query := fmt.Sprintf("Select * from %s where name=? and password=?",UserTable)
	rows:= r.db.QueryRow(query,name,password)

	err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Role)
	if err != nil {
		logrus.Errorf("could not get row from db: %s",err.Error())
		return user,err
	}

	return user,nil
}