package controllers

import (
	"errors"
	"goblog/database"
	"goblog/models"

	// "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleViewedPeoples(c *gin.Context) {
	remoteip := c.ClientIP()
	if remoteip == "::1" {
		remoteip = "127.0.0.1"
	}

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
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到记录
		vplog := models.ViwedPeoples{
			IPAddress: ip,
		}

		err = database.DB.Create(&vplog).Error
		if err != nil {
			return err
		}
		return err
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) { // 其他报错
		return err
	}

	return nil
}

// 根据ip查询记录
func FindViewedLogByIP(ip string) error {
	// log.Println(ip)
	vplog := models.ViwedPeoples{IPAddress: ip}
	result := database.DB.First(&vplog)

	if result.Error != nil {
		// log.Println(result.Error)
		return result.Error
	}

	return nil
}
