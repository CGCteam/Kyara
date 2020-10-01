package router

import (
	"server/api/v1"
	"server/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysDictionaryDetailRouter(Router *gin.RouterGroup) {
	SysDictionaryDetailRouter := Router.Group("sysDictionaryDetail").
		Use(middleware.JWTAuth()).
		Use(middleware.CasbinHandler()).
		Use(middleware.OperationRecord())
	{
		SysDictionaryDetailRouter.POST("createSysDictionaryDetail", v1.CreateSysDictionaryDetail)   // 新建SysDictionaryDetail
		SysDictionaryDetailRouter.DELETE("deleteSysDictionaryDetail", v1.DeleteSysDictionaryDetail) // 删除SysDictionaryDetail
		SysDictionaryDetailRouter.PUT("updateSysDictionaryDetail", v1.UpdateSysDictionaryDetail)    // 更新SysDictionaryDetail
		SysDictionaryDetailRouter.GET("findSysDictionaryDetail", v1.FindSysDictionaryDetail)        // 根据ID获取SysDictionaryDetail
		SysDictionaryDetailRouter.GET("getSysDictionaryDetailList", v1.GetSysDictionaryDetailList)  // 获取SysDictionaryDetail列表
	}
}
