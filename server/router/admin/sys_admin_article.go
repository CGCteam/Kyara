package admin

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1/admin"
)

func InitAdminArticle(Router *gin.RouterGroup)(R gin.IRouter)  {
	AdminArticleRouter:=Router.Group("admin")
	{
		AdminArticleRouter.POST("add",v1.AddArticle) // 添加文章
		AdminArticleRouter.GET("articleDetail",v1.GetArticleDetail)//查询文章详情
		AdminArticleRouter.POST("addNotice",v1.AddNotice)//增加公告
		AdminArticleRouter.POST("blogComment",v1.GetBlogComment)//获取博客留言
		AdminArticleRouter.POST("checkComment",v1.CheckBlogComment)//审核留言
	}
	return AdminArticleRouter
}