package blog

import (
	"server/global/response"
	"server/model"
	"server/model/request"
	resp "server/model/response"
	"server/service/blog"
	"server/utils"

	"github.com/gin-gonic/gin"
)

//博客留言
func Comment(c *gin.Context) {
	var R model.BlogComment
	_ = c.ShouldBindJSON(&R)
	if R.ID == 0 {
		ApiVerify := utils.Rules{
			"UserName":       {utils.NotEmpty()},
			"Email":          {utils.NotEmpty()},
			"Avatar":         {utils.NotEmpty()},
			"CommentContent": {utils.NotEmpty()},
		}
		ApiVerifyErr := utils.Verify(R, ApiVerify)
		if ApiVerifyErr != nil {
			response.FailWithMessage(ApiVerifyErr.Error(), c)
			return
		}
	}

	err, msg := blog.Comment(R)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkWithMessage(msg, c)
	}
}

//获取留言
func GetComment(c *gin.Context) {
	var R request.ListStruct
	_ = c.ShouldBindJSON(&R)
	err, msg, list, count := blog.GetComment(R)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(&resp.CommentMsg{Total: count, Data: list}, msg, c)
	}
}

//百度收录
func ToBaiDu(c *gin.Context) {
	var R request.ToBaiDuRequest
	err := c.ShouldBindJSON(&R)
	if err != nil {
		response.FailWithMessage("参数获取错误", c)
	}
	err, msg, count := blog.ToBaidu(R)
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(count, msg, c)
	}
}

//获取分类数量
func GetClassifyStatistics(c *gin.Context) {
	err, data, arr := blog.GetClassifyStatistics()
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkDetailed(&resp.ClassIfyRequest{Data: data, List: arr}, "获取成功", c)
	}
}

// GetBlogDynamic 获取动态
func GetBlogDynamic(c *gin.Context) {
	list, today, thatday, max, err := blog.GetBlogDynamic()
	if err != nil {
		response.FailWithMessage("获取错误", c)
	} else {
		response.OkDetailed(&request.DynamicRequest{List: list, Today: today, ThatDay: thatday, MaxCount: max}, "获取成功", c)
	}
}

//博客归档
func BlogArchives(c *gin.Context){
	err,msg,data:=blog.BlogArchives()
	if err != nil {
		response.FailWithMessage(msg, c)
	} else {
		response.OkDetailed(data,msg,c)
	}
}
