package router

import (
	"github.com/gin-gonic/gin"
	"server/api/v1"
	"server/middleware"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	AutoCodeRouter := Router.Group("autoCode").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		AutoCodeRouter.POST("createTemp", v1.CreateTemp) // 创建自动化代码
		AutoCodeRouter.GET("getTables", v1.GetTables)    // 获取对应数据库的表
		AutoCodeRouter.GET("getDB", v1.GetDB)            // 获取数据库
		AutoCodeRouter.GET("getColume", v1.GetColume)    // 获取指定表所有字段信息
	}
}
