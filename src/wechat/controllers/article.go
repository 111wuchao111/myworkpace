package controllers

import (
	"wechat/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

type Result struct {
	Posts []*models.Article `json:"posts"`
	Meta  `json:"meta"`
}

type Pagination struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Pages int   `json:"pages"`
	Total int64 `json:"total"`
	Next  bool  `json:"next"`
	Prev  bool  `json:"prev"`
}

type Meta struct {
	Pagination `json:"pagination"`
}

//获取首页文章列表
func (this *ArticleController) IndexList() {
	var articles []*models.Article
	var offset int
	var result Result
	limit := 10
	o := orm.NewOrm()
	page, _ := this.GetInt("page")
	if page > 0 {
		offset = (page - 1) * limit
	}
	o.QueryTable("article").RelatedSel().OrderBy("-create_time").Offset(offset).Limit(limit).All(&articles)
	result.Posts = articles
	total, _ := o.QueryTable("article").Count()
	result.Meta.Limit = limit
	result.Meta.Page = page
	result.Meta.Total = total
	if (total - int64(page*limit)) > 0 {
		result.Meta.Next = true
	}
	if page > 1 {
		result.Meta.Prev = true
	}

	this.Data["json"] = map[string]interface{}{"success": 0, "posts": result.Posts, "meta": result.Meta}
	this.ServeJSON()
	return
}

//专题列表
func (this *ArticleController) SubjectList() {
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "22222"}
	this.ServeJSON()
	return
}
