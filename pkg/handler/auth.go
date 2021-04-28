package handler

import (
	"fmt"
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"net/http"
)

func(h *Handler) signUp(c echo.Context) error{
	var user tictactoe_web.User

	if err := c.Bind(&user); err != nil{
		logrus.Errorf("could not bind struct User: %s",err.Error())
		return c.String(http.StatusBadGateway,"could not bind struct")
	}

	id, err := h.service.CreateUser(user)
	if err != nil{
		return c.String(http.StatusBadGateway,"could not create user")
	}

	str := fmt.Sprintf("id of created user: %d",id)
	return c.String(http.StatusOK,str)
}


func(h *Handler) signIn(c echo.Context) error{
	return nil
}
