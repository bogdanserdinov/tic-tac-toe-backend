package service

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type tokenClaims struct {
	jwt.Claims
	UserID int
}

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) CreateUser(user tictactoe_web.User) (int, error) {
	user.Password = HashPassword(user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) GenerateToken(name, password string) (string, error) {
	user,err := a.repo.GetUser(name,HashPassword(password))
	if err != nil{
		logrus.Errorf("could not get user: %s", err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12*time.Hour).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(os.Getenv("SIGNING_KEY")))
}

func (a *AuthService) ParseToken(accessToke string) (int,error){
	return 0,nil
}

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("could not bcrypt password: %s", err.Error())
	}
	return string(hashedPassword)
}
