package routers

import (
	"wechat/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ArticleController{}, "get:IndexList")
	beego.Router("/detail/?:id", &controllers.ArticleController{}, "get:Detail")
	beego.Router("/comment/?:id", &controllers.CommentController{}, "get:Get")
	beego.Router("/submitComment/?:id", &controllers.CommentController{}, "post:Post")
}
