package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleViewedPeoples(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "hello! viwedPeoples",
	})
}

// 增加一条访问记录
func CreateViewLog(ip string) {

}

// 根据ip查询记录
func FindLogByIP(ip string) {

}
