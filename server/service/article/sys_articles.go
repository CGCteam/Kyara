package service

import (
	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"server/global"
	"server/model"
	"server/model/request"
	"server/model/response"
	"strconv"
)

//文章列表服务
func ArticleList(u request.ArticleListStruct) (err error, list []model.SysArticle, totalCount int64, msg string) {
	db := global.GVA_DB
	var articleList []model.SysArticle
	var topArticle []model.SysArticle
	var total int64
	//查询总数
	err = db.Table("sys_articles").Count(&total).Error

	// 先查询置顶文章
	if u.ClassificationID == "" && u.Tag == "" {
		err = db.Where(&model.SysArticle{Top: "1"}).Find(&topArticle).Error
	}
	if u.ClassificationID != "" {
		err = db.Where(&model.SysArticle{Top: "1", ClassificationID: u.ClassificationID}).Order("updated_at desc").Find(&topArticle).Error
	}
	if u.Tag != "" {
		topArticle, err = tagToArticle(u.Tag, "1")
	}
	topLen := len(topArticle)
	//再查其他
	u.PageSize = u.PageSize - topLen
	limit := u.PageSize
	offset := u.PageSize * (u.Page - 1)
	if u.ClassificationID == "" && u.Tag == "" {
		err = db.Where("top=? ", "0").Limit(limit).Offset(offset).Order("updated_at desc").Find(&articleList).Error
	}
	if u.ClassificationID != "" {
		err = db.Where("top=? AND classification_id=?", "0", u.ClassificationID).Limit(limit).Offset(offset).Order("updated_at desc").Find(&articleList).Error
	}
	if u.Tag != "" {
		// 根据tag查询文章
		articleList, err = tagToArticle(u.Tag, "0")
	}
	articleList = append(topArticle, articleList...)
	//查询作者名
	for i, k := range articleList {
		a := model.SysUser{}
		//	根据userId查询作者姓名
		err = db.Where("UUID=?", k.UserID).Find(&a).Error
		if err != nil {
			return err, articleList, 0, "获取用户错误"
		}
		articleList[i].UserName = a.NickName
	}

	return err, articleList, total, msg
}

//根据tag_id 查询文章
func tagToArticle(tag string, top string) (articleList []model.SysArticle, err error) {
	db := global.GVA_DB
	//	先去查询中间表
	var tagLink []model.SysArticleTag
	err = db.Where("tag_id=?", tag).Find(&tagLink).Error
	// 再根据结果查询文章
	for _, i := range tagLink {
		article := model.SysArticle{}
		err = db.Where("id=? AND top=?", i.ArticleID, top).Find(&article).Error
		if err == nil {
			articleList = append(articleList, article)
		}
	}
	return articleList, err
}

//文章详情服务
func GetArticleDetail(c *gin.Context) (err error, article model.SysArticle) {
	db := global.GVA_DB
	id := c.Query("id")
	article = model.SysArticle{}
	err = db.Where("id=? ", id).Find(&article).Error
	luteEngine := lute.New()
	article.ArticleHtml = luteEngine.MarkdownStr("UvDream", article.ArticleContent)
	a := model.SysUser{}
	err = db.Where("UUID=?", article.UserID).Find(&a).Error
	article.UserName = a.NickName
	return err, article
}

//验证密码
func CheckPassword(id string, password string) (error bool, msg string) {
	article := model.SysArticle{}
	db := global.GVA_DB
	err := db.Where("id=? ", id).Find(&article).Error
	if article.ViewPassword != password && err == nil {
		return false, "密码错误"
	}
	return true, "密码正确"
}

//查询分类
func ArticleClassify() (err error, a []model.SysClassify) {
	db := global.GVA_DB
	//查询一级
	err = db.Where("parent_id=? ", "").Find(&a).Error
	if err != nil {
		return
	}
	for k, i := range a {
		a[k].Children = findChildren(i)
	}
	return err, a
}

//递归寻找分类
func findChildren(parent model.SysClassify) (children []model.SysClassify) {
	db := global.GVA_DB
	var c []model.SysClassify
	err := db.Where("parent_id=? ", parent.ID).Find(&c).Error
	if err != nil {
		return
	}
	for k, i := range c {
		c[k].Children = findChildren(i)
	}
	return c
}

//热门文章
func HotArticle() (err error, list []model.SysArticle) {
	db := global.GVA_DB
	var arr []model.SysArticle
	err = db.Order("like_count desc").Find(&arr).Error
	for k, i := range arr {
		if k < 5 {
			i.ArticleContent = ""
			list = append(list, i)
		}
	}
	return err, list
}

//所有tag
func AllTag() (err error, tagArr []model.SysTag, msg string) {
	db := global.GVA_DB
	err = db.Find(&tagArr).Error
	if err != nil {
		return err, tagArr, "获取tag失败"
	}
	for i, k := range tagArr {
		//	用tagID查询文章tag中间表
		var total int64
		err = db.Where("tag_id=?", k.ID).Table("sys_article_tags").Count(&total).Error
		if err != nil {
			return err, tagArr, "查询tag数量错误"
		}
		tagArr[i].TagCount = total
	}
	return err, tagArr, "获取tag成功"
}

//获取博客配置
func GetConfig() (err error, res response.SysConfigsResponse,msg string) {
	db := global.GVA_DB
	config := model.SysConfig{}
	err = db.Find(&config).Error
	if err!=nil{
		return err,res,"获取配置失败"
	}
	res.AuthorAvatar = config.AuthorAvatar
	//时间
	res.ActiveTime = config.ActiveTime
	res.BlogStartTime = config.BlogStartTime
	res.AuthorLink = config.AuthorLink
	res.AuthorMotto = config.AuthorMotto
	res.AuthorName = config.AuthorName
	res.BlogLogo = config.BlogLogo
	res.BlogName = config.BlogName
	//查询博客公告信息
	blogNotice:=model.BlogNotice{}
	err=db.Where("id=?",config.BlogNoticeID).Find(&blogNotice).Error
	if err!=nil{
		return err,res,"获取博客公告失败"
	}
	res.BlogNotice = blogNotice.Title
	res.BlogViewCount = config.BlogViewCount
	res.FilingMsg=config.FilingMsg
	//查询文章数量
	err=db.Table("sys_articles").Count(&res.ArticleCount).Error
	if err!=nil{
		return err,res,"获取文章数量失败"
	}
	return err, res,"获取配置成功"
}
//获取github仓库列表
func GetGithub()(githubList []response.GithubList,err error)  {
	blogConfig:= model.SysConfig{}
	db := global.GVA_DB
	err=db.Find(&blogConfig).Error
	r,err:=req.Get("https://api.github.com/users/"+blogConfig.GithubUserName+"/repos")
	if err==nil{
		r.ToJSON(&githubList)
	}
	return githubList,err
}
//访问博客
func ViewBlogCount(c *gin.Context)(err error,msg string)  {
	db := global.GVA_DB
	id := c.Query("id")
	if id!="blog"{
		article:=model.SysArticle{}
		err=db.Where("id=?",id).Find(&article).Error
		if err==nil{
			num, err:= strconv.ParseInt(article.ViewCount, 10, 64)
			num=num+1
			article.ViewCount=strconv.FormatInt(num, 10)
			err=db.Model(&article).Where("id=?",id).Update("view_count",article.ViewCount).Error
			if err!=nil{
				return err,"更新文章访问量失败"
			}
			return err,"更新文章访问量成功"
		}
	}
	if id=="blog"{
		config:=model.SysConfig{}
		err=db.Find(&config).Error
		num, err:= strconv.ParseInt(config.BlogViewCount, 10, 64)
		num=num+1
		config.BlogViewCount=strconv.FormatInt(num, 10)
		err=db.Model(&config).Update("blog_view_count",config.BlogViewCount).Error
		if err!=nil{
			return err,"更新博客访问量失败"
		}
		return err,"更新博客访问量成功"
	}
	return err,"更新访问量成功"
}
//获取公告
func GetNotice(c *gin.Context)(err error,list model.BlogNotice,msg string)  {
	db := global.GVA_DB
	err=db.Find(&list).Error
	if err!=nil{
		return err,list,"获取公告失败"
	}
	return  err,list,"获取公告成功"
}