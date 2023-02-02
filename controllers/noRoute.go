package controllers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func HandleNoRoute(c *gin.Context) {
	u := url.URL{Path: "/home"}
	c.Redirect(http.StatusFound, u.RequestURI())
}
