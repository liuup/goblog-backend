package controllers

import (
	"errors"
	"goblog/database"
	"goblog/models"
	"log"

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
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"viewed_peoples": GetViewedPeoplesCount(),
	})
}

// 增加一条独立访问记录
func CreateViewLog(ip string) error {
	err := FindViewedLogByIP(ip)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到记录
		vplog := models.ViewedPeoples{
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
	vplog := models.ViewedPeoples{IPAddress: ip}
	result := database.DB.First(&vplog)

	if result.Error != nil {
		// log.Println(result.Error)
		return result.Error
	}

	return nil
}

// // TODO: 获取网站独立访问量的人数
func GetViewedPeoplesCount() int64 {
	var count int64

	// 查找没有没删除的访问总数
	result := database.DB.Table("viewed_peoples").Where("is_deleted = ?", 0).Count(&count)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}

	return count
}
