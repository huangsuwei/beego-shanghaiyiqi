package controllers

import (
	"beego-shanghaiyiqi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (userC *UserController) ShowRegister() {
	userC.TplName = "register.html"
}

func (userC *UserController) Register() {
	//获取数据
	name := userC.GetString("name")
	password := userC.GetString("password")
	//beego.Info(name, password)
	//校验数据
	userC.TplName = "register.html"
	if name == "" || password == "" {
		userC.Data["errmsg"] = "用户名或密码不得为空"
		return
	}
	//插入数据
	o := orm.NewOrm()
	var user models.User
	user.Name = name
	err := o.Read(&user, "name")
	if err != nil {
		userC.Data["errmsg"] = err
	}
	if user.Id != 0 {
		userC.Data["errmsg"] = "用户名重复！"
		return
	}
	/*user.Name = name
	err := o.Read(&user, "name")
	if err != nil {
		userC.Data["errmsg"] = err
		return
	}
	if user.Id != 0 {
		userC.Data["errmsg"] = "用户名重复！"
		return
	}*/
	//返回
	user.Name = name
	user.Password = password
	_, err1 := o.Insert(&user)
	if err1 != nil {
		userC.Data["errmsg"] = err1
		return
	}
	//userC.Data["errmsg"] = "写入成功！"
	//userC.Ctx.WriteString("注册成功！！")
	userC.Redirect("/login", 302)
	//换页面也是可行的，但是url是不会变的，所以用上面这个更好
	//userC.TplName = "login.html"
}

func (userC *UserController) ShowLogin() {
	userC.TplName = "login.html"
	userName := userC.Ctx.GetCookie("userName")
	userC.Data["userName"] = userName
	userC.Data["checked"] = ""
	if userName == "" {
		userC.Data["checked"] = "checked"
	}
}

func (userC *UserController) Login() {
	userC.TplName = "login.html"
	//获取数据
	name := userC.GetString("name")
	password := userC.GetString("password")
	remember := userC.GetString("remember")
	if name == "" || password == "" {
		userC.Data["errmsg"] = "用户名或密码不得为空"
		return
	}
	o := orm.NewOrm()
	var user models.User
	user.Name = name
	err := o.Read(&user, "name")
	if err != nil {
		userC.Data["errmsg"] = "用户不存在!"
		return
	}
	if user.Password != password {
		userC.Data["errmsg"] = "密码错误!"
		return
	}

	//userC.Ctx.WriteString("登陆成功！！")
	//设置cookie
	if remember == "on" {
		userC.Ctx.SetCookie("userName", name, 100)
	} else {
		userC.Ctx.SetCookie("userName", name, -1)
	}

	userC.SetSession("userName", name)
	userC.Redirect("/ArticleList", 302)
}

//退出登陆
func (userC *UserController) Logout() {
	//删除session
	userC.DelSession("userName")
	//跳转到登陆页面，
	userC.Redirect("/login", 302)
}

func (userC *UserController) GetUserId() int {
	name := userC.GetSession("userName")
	if name == nil {
		userC.Redirect("/login", 302)
		return 0
	}
	o := orm.NewOrm()
	var user models.User
	user.Name = name.(string)
	o.Read(&user, "name")

	return user.Id
}

func (userC *UserController) GetUser() models.User {
	name := userC.GetSession("userName")
	beego.Info("name:", name)
	var user models.User
	if name == nil {
		userC.Redirect("/login", 302)
	}
	o := orm.NewOrm()
	user.Name = name.(string)
	o.Read(&user, "name")

	return user
}
