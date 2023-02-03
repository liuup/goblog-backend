package controllers

import (
	"goblog/database"
	"goblog/models"
	"log"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleViewedPeoples(c *gin.Context) {
	remoteip := c.ClientIP()
	// if remoteip == "::1" {
	// 	remoteip = "127.0.0.1"
	// }

	// 尝试根据ip创建记录
	err := CreateViewLog(remoteip)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "hello! viwedPeoples",
	})
}

// 增加一条访问记录
func CreateViewLog(ip string) error {
	err := FindViewedLogByIP(ip)
	if err != nil { // 如果没找记录
		return err
	}

	// 如果没找到记录
	vplog := models.ViwedPeoples{
		IPAddress: ip,
	}

	err = database.DB.Create(&vplog).Error
	if err != nil {
		return err
	}

	return nil
}

// 根据ip查询记录
func FindViewedLogByIP(ip string) error {
	// log.Println(ip)
	vplog := models.ViwedPeoples{}
	result := database.DB.Find(&vplog, "ip_address = ?", ip)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
