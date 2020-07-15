package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["data"] = "china,chinese"
	//c.TplName = "index.tpl"
	c.TplName = "test.html"
}

func (c *MainController) Post() {
	c.Data["data"] = "go语言天下无敌"
	//c.TplName = "index.tpl"
	c.TplName = "test.html"
}
