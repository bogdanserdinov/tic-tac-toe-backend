package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const userCtx = "userId"

func (h *Handler) CheckUser(c *gin.Context) {

	header := c.GetHeader("Authorization")

	if header == "" {
		c.String(http.StatusUnauthorized, "empty auth header")
		return
	}

	headerPath := strings.Split(header, " ")

	if len(headerPath) != 2 || headerPath[0] != "Bearer" {
		c.String(http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerPath[1]) == 0 {
		c.String(http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerPath[1])
	if err != nil {
		message := fmt.Sprintf("invalid token: %s", err.Error())
		c.String(http.StatusUnauthorized, message)
		return
	}

	c.Set(userCtx, userId)

}

func GetUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("no user id in context")
	}
	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("invalid type of id")
	}
	return idInt, nil
}
