package handler

import (
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/service"
	"github.com/labstack/echo"
)

type Handler struct{
	service *service.Service
}

func NewHandler(service *service.Service) *Handler{
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *echo.Echo{
	e := echo.New()

	auth := e.Group("/auth")
	auth.POST("/sign-in",h.signIn)
	auth.POST("/sign-up",h.signUp)

	return e
}