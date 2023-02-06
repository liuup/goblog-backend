package controllers

import (
	"goblog/database"
	"goblog/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用cron定时任务来获取服务器负载信息
func HandleServerStatusCron(c *gin.Context) {
	ss := models.ServerStatus{}

	result := database.DB.Last(&ss)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ss,
	})
}
