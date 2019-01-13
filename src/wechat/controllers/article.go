package controllers

import (
	"wechat/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

//获取首页文章列表
func (this *ArticleController) IndexList() {
	var articles []*models.Article
	var offset int
	limit := 2
	o := orm.NewOrm()
	page, _ := this.GetInt("page")
	if page > 0 {
		offset = (page - 1) * limit
	}
	o.QueryTable("article").RelatedSel().OrderBy("-create_time").Offset(offset).Limit(limit).All(&articles)
	this.Data["json"] = map[string]interface{}{"success": 0, "data": articles}
	this.ServeJSON()
	return
}

//专题列表
func (this *ArticleController) SubjectList() {
	this.Data["json"] = map[string]interface{}{"success": 0, "message": "22222"}
	this.ServeJSON()
	return
}
