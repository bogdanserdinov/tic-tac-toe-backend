package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) GetStats(c echo.Context) error{
	id := c.Get("userID")

	res := fmt.Sprintf("id : %d",id)
	return c.String(http.StatusOK,res)
}
