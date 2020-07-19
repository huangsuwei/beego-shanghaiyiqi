package controllers

import (
	"beego-shanghaiyiqi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

//展示文章列表页面
func (ac *ArticleController) ArticleList() {
	ac.TplName = "index.html"
	//获取数据展示
	o := orm.NewOrm()
	at := o.QueryTable("Article")
	var articles []models.Article
	//查询数量
	totalCount, err := at.All(&articles)
	pageSize := 2
	pageIndex, err2 := ac.GetInt("pageIndex")
	if err2 != nil {
		//没传数据，访问首页
		pageIndex = 0
	}
	//起始位置
	at.Limit(pageSize, pageSize*(pageIndex-1)).RelatedSel("ArticleType").All(&articles)

	pageCount := math.Ceil(float64(totalCount) / float64(pageSize))
	if err != nil {
		beego.Info("查询错误！")
	}
	//查询总记录数
	totalCount, err1 := at.Count()
	if err1 != nil {
		beego.Info("查询错误！")
	}
	ac.Data["articles"] = articles
	ac.Data["totalCount"] = totalCount
	ac.Data["pageCount"] = int(pageCount)
	if pageIndex == 0 {
		pageIndex = 1
	}
	ac.Data["pageIndex"] = pageIndex
	ac.Data["types"] = new(ArticleTypeController).GetAllTypes()
}

func (ac *ArticleController) ShowAddArticle() {
	ac.TplName = "add.html"
	typeC := new(ArticleTypeController)
	ac.Data["types"] = typeC.GetAllTypes()
}

func (ac *ArticleController) AddArticle() {
	beego.Info("id-2:")
	ac.TplName = "add.html"
	id, _ := ac.GetInt("id")
	typeName := ac.GetString("select")
	articleName := ac.GetString("articleName")
	content := ac.GetString("content")
	if articleName == "" || content == "" {
		ac.Data["errmsg"] = "文章标题和内容不得为空"
		return
	}
	beego.Info("id-1:", id)

	//图片没有上传时，可以空
	_, head, err := ac.GetFile("uploadname")
	//defer file.Close()//不关闭报错
	if err != nil {
		ac.Data["errmsg"] = err
		return
	}
	if head.Size > 5000000 {
		ac.Data["errmsg"] = "文件过大，请重新上传"
		return
	}

	ext := path.Ext(head.Filename)
	exts := []string{".jpg", ".png", ".jpeg"}
	inExt := false
	for _, strictExt := range exts {
		if strictExt == ext {
			inExt = true
			break
		}
	}
	if !inExt {
		ac.Data["errmsg"] = "文件格式仅允许\".jpg\", \".png\", \".jpeg\"，请重新上传"
		return
	}
	beego.Info("id0:", id)

	//文件重名
	fileName := time.Now().Format("2006-01-02 15:04:05") + ext
	//这里为什么要加一个点？框架本身问题，
	err = ac.SaveToFile("uploadname", "./static/img/"+fileName)
	if err != nil {
		ac.Data["errmsg"] = err
		return
	}
	beego.Info("id1:", id)

	o := orm.NewOrm()
	var article models.Article
	if id != 0 {
		article.Id = id
		err3 := o.Read(&article)
		if err3 != nil {
			ac.Data["errmsg"] = "找不到对应的文章！"
			return
		}
	}
	article.ArtiName = articleName
	article.Acontent = content
	article.Aimg = "/static/img/" + fileName
	var articleType models.ArticleType
	articleType.TypeName = typeName
	err5 := o.Read(&articleType, "type_name")
	if err5 != nil {
		ac.Data["errmsg"] = err5
		return
	}
	article.ArticleType = &articleType
	beego.Info("id2:", id)
	if id != 0 {
		updateCount, err4 := o.Update(&article)
		beego.Info(err4)
		if err4 != nil {
			ac.Data["errmsg"] = err4
			return
		}
		ac.Data["updateCount"] = updateCount
	} else {
		o.Insert(&article)
	}

	ac.Redirect("/ArticleList", 302)
}

//文章详情
func (ac *ArticleController) ArticleDetail() {
	ac.TplName = "content.html"
	id, _ := ac.GetInt("id")
	isEdit, _ := ac.GetInt("is_edit")
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	err := o.Read(&article)
	if err != nil {
		ac.Data["errmsg"] = err
	}
	//增加阅读次数
	article.Acount += 1
	_, err2 := o.Update(&article)
	if err2 != nil {
		ac.Data["errmsg"] = err2
	}
	if isEdit == 1 {
		ac.Data["id"] = id
		ac.TplName = "update.html"
	}
	ac.Data["article"] = article
}

func (ac *ArticleController) ArticleDel() {
	id, _ := ac.GetInt("id")
	o := orm.NewOrm()
	var article models.Article
	article.Id = id
	err := o.Read(&article)
	if err != nil {
		ac.Data["errmsg"] = err
	}
	_, err2 := o.Delete(&article)
	if err2 != nil {
		ac.Data["errmsg"] = err2
	}
	ac.Redirect("/ArticleList", 302)
}
