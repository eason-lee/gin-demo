package router

import (
	_ "gin-demo/docs"
	"gin-demo/api/v1"
	"gin-demo/global"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	
)

func Run() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.LOG.Info("register swagger handler")
	
	// 简单的路由组: v1
	statisticsRouter := router.Group("v1/statistics")
	{
		statisticsRouter.POST("meta", v1.CreateMetas)
		statisticsRouter.GET("meta", v1.GetMetas)
	}
	router.Run()
}
