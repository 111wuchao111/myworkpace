package controllers

import (
	"wechat/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CommentController struct {
	beego.Controller
}

//获取某个文章的评论
func (this *CommentController) Get() {
	var Comments []*models.Comments
	blogId := this.Ctx.Input.Param(":id")
	o := orm.NewOrm()
	o.QueryTable("article_comments").Filter("article_id", blogId).OrderBy("-create_time").All(&Comments)
	this.Data["json"] = map[string]interface{}{"success": 0, "data": Comments}
	this.ServeJSON()
}

//提交某个文章的评论
func (this *CommentController) Post() {
	this.Data["json"] = map[string]interface{}{"success": 0, "data": this.Ctx.Input.RequestBody}
	this.ServeJSON()
}
