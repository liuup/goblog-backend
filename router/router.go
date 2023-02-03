package router

import (
	"goblog/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// 公共api
	r.GET("/home", controllers.HandleHome)            // 主页
	r.GET("/about")                                   // 关于本人
	r.GET("/viewed", controllers.HandleViewedPeoples) // 处理访问人数

	r.GET("/article", controllers.HandleArticle)          // 获取特定文章内容
	r.GET("/comments", controllers.HandleArticleComments) // 获取文章评论

	r.GET("/server_status", controllers.HandleServerStatus)          // TODO: 实时获取服务器负载信息
	r.GET("/server_status_cron", controllers.HandleServerStatusCron) // TODO: 使用cron定时任务获取服务器负载信息

	// r.NoRoute(controllers.HandleNoRoute) // 404路由

	// 后台私有api
	admin := r.Group("/v1")
	{
		admin.POST("/login")
	}

	return r
}
