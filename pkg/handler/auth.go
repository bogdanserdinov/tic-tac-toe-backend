package handler

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

func(h *Handler) signIn(c echo.Context) error{
	var user tictactoe_web.User

	if err := c.Bind(user); err != nil{
		logrus.Errorf("could not bind struct User: %s",err.Error())
		return c.String(http.StatusBadRequest,"could not bind struct")
	}

	return nil
}


func(h *Handler) signUp(c echo.Context) error{
	return nil
}
