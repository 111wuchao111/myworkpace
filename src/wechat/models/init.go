package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(ArticleClass), new(Article), new(User))
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:@/wechatapp?charset=utf8")
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
}
