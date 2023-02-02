package router

import (
	"goblog/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 公共api
	r.GET("/home", controllers.HandleHome) // 主页
	// r.GET("/about")  // 关于本人
	r.GET("/viewed", controllers.HandleViewedPeoples) // 处理访问人数

	// 404路由
	r.NoRoute(controllers.HandleNoRoute)

	// 私有api

	return r
}
