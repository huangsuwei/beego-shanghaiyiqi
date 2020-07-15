package routers

import (
	"beego-shanghaiyiqi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//给请求自定义方法,指定一个方法
	/*beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin")
	    //多个请求指定一个方法
		beego.Router("/loginOut", &controllers.LoginController{}, "get,post:ShowLogin")
	    //所有指定
		beego.Router("/loginIndex", &controllers.LoginController{}, "*:ShowLogin")
		//同时指定时，会优先访问指定了方法的方法
		beego.Router("/loginTest", &controllers.LoginController{}, "*:ShowLogin;get:show")*/
}
