package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "home",
	})
}
