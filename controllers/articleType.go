package controllers

import (
	"beego-shanghaiyiqi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ArticleTypeController struct {
	beego.Controller
}

func (at *ArticleTypeController) ShowAddType() {
	at.TplName = "addType.html"
	at.Data["types"] = at.GetAllTypes()
}

func (at *ArticleTypeController) ShowAllTypes() {
	at.TplName = "addType.html"
	at.Data["types"] = at.GetAllTypes()
}

func (at *ArticleTypeController) GetAllTypes() []models.ArticleType {
	o := orm.NewOrm()
	var types []models.ArticleType
	o.QueryTable("ArticleType").All(&types)

	return types
}

func (at *ArticleTypeController) AddType() {
	at.TplName = "add.html"
	typeName := at.GetString("typeName")
	if typeName == "" {
		at.Data["errmsg"] = "分类名不得为空"
		return
	}
	var articleTYpe models.ArticleType
	articleTYpe.TypeName = typeName
	o := orm.NewOrm()
	_, err := o.Insert(&articleTYpe)
	if err != nil {
		at.Data["errmsg"] = err
		return
	}
	at.Redirect("/AddType", 302)
}

func (at *ArticleTypeController) GetIdByName(name string) int {
	o := orm.NewOrm()
	var atype models.ArticleType
	atype.TypeName = name
	err := o.Read(&atype, "type_name")
	if err != nil {
		beego.Info(err)
		return 0
	}

	return atype.Id
}
