package repository

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthRepository struct{
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository{
	return &AuthRepository{
		db : db,
	}
}

func(r *AuthRepository) CreateUser(user tictactoe_web.User) (int,error){
	//	res, err := conn.Exec("INSERT INTO users (name) VALUES(\"Peter\")")
	var id int64
	res, err := r.db.Exec("INSERT INTO user (name,password) values ($1,$2)",user.Name,user.Password)
	if err != nil{
		logrus.Fatalf( "could execute insert query: %s",err.Error())
		return 0,err
	}
	//create a record in stats table with id of user
	id,_ = res.LastInsertId()
	_, err = r.db.Exec("INSERT INTO userstats (id,total,wins,draws,losses) values ($1,0,0,0,0)",id)
	if err != nil{
		logrus.Fatalf( "could execute insert query: %s",err.Error())
		return 0,err
	}

	return int(id),nil
}

func(r *AuthRepository) GetUser(name string, password string) (tictactoe_web.User,error){
	var user tictactoe_web.User
	err := r.db.Get(&user, "SELECT * FROM user WHERE name=$1 and password=$2", name,password)

	if err != nil {
		logrus.Errorf("could not get row from db: %s",err.Error())
		return user,err
	}

	return user,nil
}