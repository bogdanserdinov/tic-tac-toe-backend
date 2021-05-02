package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) GetStats(c echo.Context) error{
	id,err := GetUserId(c)
	fmt.Println("dbsjkfk")
	if err != nil {
		return c.String(http.StatusInternalServerError,"could not get user id")
	}
	res := fmt.Sprintf("id : %d",id)
	return c.String(http.StatusOK,res)
}
