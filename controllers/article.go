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
	if err != nil {
		beego.Info("查询错误！")
	}
	//查询总记录数
	totalCount, err1 := at.Count()
	if err1 != nil {
		beego.Info("查询错误！")
	}
	pageSize := 2
	pageCount := math.Ceil(float64(totalCount) / float64(pageSize))
	ac.Data["articles"] = articles
	ac.Data["totalCount"] = totalCount
	ac.Data["pageCount"] = pageCount
}

func (ac *ArticleController) ShowAddArticle() {
	ac.TplName = "add.html"
}

func (ac *ArticleController) AddArticle() {
	ac.TplName = "add.html"
	articleName := ac.GetString("articleName")
	content := ac.GetString("content")
	if articleName == "" || content == "" {
		ac.Data["errmsg"] = "文章标题和内容不得为空"
		return
	}

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

	//文件重名
	fileName := time.Now().Format("2006-01-02 15:04:05") + ext
	//这里为什么要加一个点？框架本身问题，
	err = ac.SaveToFile("uploadname", "./static/img/"+fileName)
	if err != nil {
		ac.Data["errmsg"] = err
		return
	}

	o := orm.NewOrm()
	var article models.Article
	article.ArtiName = articleName
	article.Acontent = content
	article.Aimg = "/static/img/" + fileName
	o.Insert(&article)

	ac.Redirect("/ArticleList", 302)
}
