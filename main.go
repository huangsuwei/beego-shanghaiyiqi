package main

import (
	_ "beego-shanghaiyiqi/models"
	_ "beego-shanghaiyiqi/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.AddFuncMap("prePage", ShowPrePage)
	beego.AddFuncMap("nextPage", ShowNextPage)
	beego.Run()
}

//后台定义函数
//视图函数，在beego.run前，处理
func ShowPrePage(pageIndex int) int {
	if pageIndex <= 1 {
		return 1
	}

	return pageIndex - 1
}

func ShowNextPage(pageIndex, pageCount int) int {
	nextPage := pageIndex + 1
	if nextPage >= pageCount {
		return pageCount
	}

	return nextPage
}
