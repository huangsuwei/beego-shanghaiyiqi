package controllers

import (
	"beego-shanghaiyiqi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

func (c *MainController) ShowGet() {
	//获取orm操作对象
	o := orm.NewOrm()
	//执行相应的操作，增删改查
	//插入
	var user models.User
	user.Name = "huangsuwei"
	user.Password = "123456"
	addCount, err1 := o.Insert(&user)
	if err1 != nil {
		beego.Info(err1)
	}
	c.Data["data"] = addCount
	//c.TplName = "index.tpl"
	c.TplName = "test.html"
}
