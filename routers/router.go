package routers

import (
	"beego-shanghaiyiqi/controllers"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/addUser1", &controllers.MainController{}, "get:ShowGet")
	//给请求自定义方法,指定一个方法
	/*beego.Router("/login", &controllers.LoginController{}, "get:ShowLogin")
	    //多个请求指定一个方法
		beego.Router("/loginOut", &controllers.LoginController{}, "get,post:ShowLogin")
	    //所有指定
		beego.Router("/loginIndex", &controllers.LoginController{}, "*:ShowLogin")
		//同时指定时，会优先访问指定了方法的方法
		beego.Router("/loginTest", &controllers.LoginController{}, "*:ShowLogin;get:show")*/
	//这个页面不重启，会找不到方法名
	beego.Router("/register", &controllers.UserController{}, "get:ShowRegister")
	beego.Router("/register", &controllers.UserController{}, "post:Register")

	//登陆
	beego.Router("/login", &controllers.UserController{}, "get:ShowLogin")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	//beego.Router("/article/Logout", &controllers.UserController{}, "get:Logout", Filter)//过滤器
	beego.Router("/Logout", &controllers.UserController{}, "get:Logout")

	//文章列表页面
	beego.Router("/ArticleList", &controllers.ArticleController{}, "get:ArticleList")

	//添加文章
	beego.Router("/AddArticle", &controllers.ArticleController{}, "get:ShowAddArticle;post:AddArticle")

	//文章详情
	beego.Router("/ArticleDetail", &controllers.ArticleController{}, "get:ArticleDetail")

	//文章删除
	beego.Router("/ArticleDel", &controllers.ArticleController{}, "get:ArticleDel")

	//文章编辑
	//beego.Router("/AddArticle", &controllers.ArticleController{}, "get:AddArticle")
	//分类添加
	beego.Router("/AddType", &controllers.ArticleTypeController{}, "get:ShowAddType;post:AddType")

	beego.Router("/ShowAllTypes", &controllers.ArticleTypeController{}, "get:ShowAllTypes")
	beego.Router("/DelType", &controllers.ArticleTypeController{}, "get:DelType")
}

//过滤器先不用
/*var Filter = func(ctx context.Context) {
	name := ctx.Input.Session("userName")
	if name == nil {
		ctx.Redirect(302, "/login")
	}
}*/
