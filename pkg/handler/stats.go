package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetStats(c *gin.Context){
	id,err := GetUserId(c)
	if err != nil {
		c.String(http.StatusInternalServerError,"could not get user id")
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"id" : id,
	})
}
