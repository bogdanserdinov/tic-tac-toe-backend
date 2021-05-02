package handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

const userCtx = "userId"


func (h *Handler) CheckUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error{

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
			message := fmt.Sprintf("invalid token: %s",err.Error())
			return c.String(http.StatusUnauthorized,message )
		}

		c.Set(userCtx, userId)
		return nil
	}
}

func GetUserId(c echo.Context) (int,error){
		id := c.Get(userCtx)

		idInt, ok := id.(int)
		if !ok {
			return 0, errors.New("invalid type of id")
		}
		fmt.Println(id)
		fmt.Println("asfknsadflsan")
		return idInt, nil
	}

