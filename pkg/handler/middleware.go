package handler

import (
	"errors"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

func (h *Handler) checkUser(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		return c.String(http.StatusUnauthorized, "empty auth header")
	}

	headerPath := strings.Split(header, " ")

	if len(headerPath) != 2 || headerPath[0] != "Bearer" {
		return c.String(http.StatusUnauthorized, "invalid auth header")
	}

	if len(headerPath[1]) == 0 {
		return c.String(http.StatusUnauthorized, "token is empty")
	}

	userId, err := h.service.Authorization.ParseToken(headerPath[1])
	if err != nil {
		return c.String(http.StatusUnauthorized, "invalid token")
	}

	c.Set("userId", userId)
	return nil
}

func (h *Handler) getUser(c echo.Context) (int, error) {
	id := c.Get("userId")

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("invalid type of id")
	}

	return idInt, nil
}
