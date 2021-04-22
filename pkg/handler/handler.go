package handler

import "github.com/labstack/echo"

type Handler struct{

}

func NewHandler() *Handler{
	return &Handler{}
}

func (h *Handler) InitRoutes() *echo.Echo{
	e := echo.New()

	auth := e.Group("/auth")
	auth.POST("/sign-in",h.signIn)
	auth.POST("/sign-up",h.signUp)

	return e
}