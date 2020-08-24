package model

import "github.com/jinzhu/gorm"

// SysArticle 文章数据表
type SysArticle struct {
	gorm.Model
	// 文章标题
	Title string `json:"title" gorm:"comment:'文章标题'"`
	// 文章摘要
	Introduction string `json:"introduction" gorm:"comment:'文章摘要'"`
	// 作者id
	UserID string `json:"user_id" gorm:"comment:'作者id'"`
	// ---作者名---
	UserName string `json:"user_name" gorm:"-"`
	// 评论id
	CommentID string `json:"comment_id" gorm:"comment:'评论id'"`
	// icon
	Icon string `json:"icon" gorm:"comment:'icon'"`
	// icon颜色
	IconColor string `json:"icon_color" gorm:"comment:'icon颜色'"`
	// 图片地址
	ImgURL string `json:"img_url" gorm:"comment:'图片地址'"`
	// 文章内容
	ArticleContent string `json:"article_content" gorm:"comment:'文章内容';type:longtext "`
	// 浏览量
	ViewCount string `json:"view_count" gorm:"comment:'浏览量'"`
	// 原文链接
	OriginalLink string `json:"original_link" gorm:"comment:'原文链接'"`
	// 标签id
	TagID string `json:"tag_id" gorm:"comment:'标签id'"`
	// 分类id
	ClassificationID string `json:"classification_id" gorm:"comment:'分类id'"`
	// 阅读密码
	ViewPassword string `json:"view_password" gorm:"comment:'阅读密码'"`
	// 是否开启评论
	IsComment string `json:"is_comment" gorm:"comment:'是否开启评论'"`
	// 评论量
	CommentCount string `json:"comment_count" gorm:"comment:'评论量'"`
	// 状态
	Status string `json:"status" gorm:"comment:'状态';default:'0'"`
	// 是否需要密码
	IsPassword string `json:"is_password" gorm:"comment:'是否需要密码';default:'0'"`
	// 转载规则设置
	TransferRules string `json:"transfer_rules" gorm:"comment:'转载规则设置';default:'0'"`
	// 封面类型
	CoverType string `json:"cover_type" gorm:"comment:'封面类型';default:'0'"`
	// 是否置顶
	Top string `json:"top" gorm:"comment:'是否置顶';default:'false'"`
}