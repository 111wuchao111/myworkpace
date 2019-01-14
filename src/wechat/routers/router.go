package routers

import (
	"wechat/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ArticleController{}, "get:IndexList")
	beego.Router("/detail/?:id", &controllers.ArticleController{}, "get:Detail")
}
