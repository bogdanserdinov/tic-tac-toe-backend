package handler

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var user tictactoe_web.User

	if err := c.Bind(&user); err != nil {
		logrus.Errorf("could not bind struct User: %s", err.Error())
		c.String(http.StatusBadRequest, "could not bind struct")
		return
	}

	id, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		c.String(http.StatusInternalServerError, "could not create user")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
	return
}

func (h *Handler) signIn(c *gin.Context) {
	var user tictactoe_web.User

	if err := c.Bind(&user); err != nil {
		logrus.Errorf("could not bind struct User: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "could not bind struct",
		})
	}

	token, err := h.service.Authorization.GenerateToken(user.Name, user.Password)
	if err != nil {
		logrus.Errorf("could not parse token: %s", err.Error())
		c.String(http.StatusInternalServerError, "could not parse token")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
	return
}
