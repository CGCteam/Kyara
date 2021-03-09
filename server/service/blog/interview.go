package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/model"
	"server/model/request"
)
//新增面试题分类
func AddInterviewClassify(r model.InterviewClassify)(err error,msg string) {
	fmt.Println(r)
	db := global.GVA_DB
	//先校验名称
	err=db.Where("classify_name=?",r.ClassifyName).Find(&r).Error
	if err==nil {
		return err,"该名称已经存在"
	}
	err=db.Create(&r).Error
	if err!=nil {
		return err,"创建分类失败"
	}
	return err,"创建分类成功"
}
//获取面试题分类
func GetInterviewClassify() (err error, msg string, data []model.InterviewClassify) {
	db := global.GVA_DB
	err = db.Find(&data).Error
	if err != nil {
		return err, "获取分类失败", data
	}
	return err, "获取分类成功", data
}

//获取面试题
func GetInterview(c *gin.Context, r request.InterviewSearch) (err error, msg string, interview []model.Interview, totalCount int64) {
	db := global.GVA_DB
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	//获取面试题总数
	err = db.Find(&interview).Count(&totalCount).Error
	if err != nil {
		return err, "获取题库", interview, 0
	}
	//获取面试题
	//if r.LevelSort!="" {
	//难度排序以及level查询
	err = db.Where("level LIKE ?", "%"+r.Level+"%").Where("classify = ?", r.Classify).Order("level " + r.LevelSort).Find(&interview).Limit(limit).Offset(offset).Error
	//}else{
	//	err=db.Where("level LIKE ?","%"+r.Level+"%").Where("classify = ?", r.Classify).Find(&interview).Limit(limit).Offset(offset).Error
	//}
	if err != nil {
		return err, "获取题库失败", interview, totalCount
	}
	return err, "获取题库成功", interview, totalCount
}
