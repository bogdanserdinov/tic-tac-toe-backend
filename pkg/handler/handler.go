package handler

import (
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	auth.POST("/sign-in", h.signIn)
	auth.POST("/sign-up", h.signUp)

	api := router.Group("/api", h.CheckUser)
	{

		stats := api.Group("/profile")
		{
			stats.GET("", h.GetStats)
			//stats.POST("/profile",)
		}

		//bot := api.Group("/bot")
		//{
		//bot.POST("/bot")
		//}

		//game := api.Group("/game")
		//{
		//game.GET("/game")
		//game.POST("/game")
		//}

	}
	return router
}
