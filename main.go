package main

import (
	_ "beego-shanghaiyiqi/models"
	_ "beego-shanghaiyiqi/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
