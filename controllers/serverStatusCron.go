package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: 使用cron定时任务来获取服务器负载信息
func HandleServerStatusCron(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"server_status_cron": "server cron ok",
	})
}
