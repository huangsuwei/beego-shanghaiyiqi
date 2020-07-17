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
	/*user.Name = "huangsuwei"
	user.Password = "123456"
	addCount, err1 := o.Insert(&user)
	c.Data["data"] = addCount
	*/
	//查询操作
	/*user.Id = 1
	//有主键，可以不要查询的字段
	err1 := o.Read(&user, "id")
	if err1 != nil {
		beego.Info(err1)
	}
	beego.Info(user)*/
	//更新操作
	/*user.Id = 2
	err1 := o.Read(&user)
	if err1 != nil {
		beego.Info("要更新的数据不存在")
	}
	user.Name = "caoxixiu"
	updateCount, err2 := o.Update(&user)
	if err2 != nil {
		beego.Info("要更新的数据不存在")
	}*/
	//删除操作
	user.Id = 1
	deleteCount, err1 := o.Delete(&user)
	if err1 != nil {
		beego.Info("要删除的数据不存在")
	}
	c.Data["data"] = deleteCount
	//c.TplName = "index.tpl"
	c.TplName = "test.html"
}
