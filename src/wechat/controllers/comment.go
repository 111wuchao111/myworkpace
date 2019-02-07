package controllers

import (
	"wechat/models"

	"time"

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
	o := orm.NewOrm()
	var Comments []*models.Comments
	var Comment models.Comments
	blogId, _ := this.GetInt("blogId")
	Comment.ArticleId = blogId
	Comment.UserOpenid = this.GetString("openId")
	Comment.UserHeadicon = this.GetString("cAvatarUrl")
	Comment.UserNickname = this.GetString("cNickName")
	Comment.Content = this.GetString("comment")
	Comment.CreateTime = time.Now().Format("2006-01-02")
	beego.Info(Comment)
	_, err := o.Insert(&Comment)
	if err == nil {
		o.QueryTable("article_comments").Filter("article_id", blogId).OrderBy("-create_time").All(&Comments)
		this.Data["json"] = map[string]interface{}{"success": 0, "data": Comments}
		this.ServeJSON()
	}
	this.Data["json"] = map[string]interface{}{"success": -1, "data": "error"}
	this.ServeJSON()
}
