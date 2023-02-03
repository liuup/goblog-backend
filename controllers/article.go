package controllers

import (
	"goblog/database"
	"goblog/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 处理文章入口函数
func HandleArticle(c *gin.Context) {
	article_id := c.Query("id") // 获取文章id

	article := models.Article{}

	result := database.DB.Table("articles").Where("article_id = ?", article_id).Where("is_deleted = ?", 0).Select("article_id", "content", "create_time").First(&article)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	log.Println(article)

	c.JSON(http.StatusOK, gin.H{
		"article": gin.H{
			"ArticleId":  article.ArticleId,
			"Content":    article.Content,
			"CreateTime": article.CreateTime,
		},
	})
}

// 获取文章总数
func GetArticleCount() int64 {
	var count int64

	result := database.DB.Table("articles").Count(&count)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return count
}
