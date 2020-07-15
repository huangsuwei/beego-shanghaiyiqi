package models

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//操作数据库代码
	conn, err := sql.Open("mysql", "root:@tcp(10.10.10.201:3306)/test?charset=utf8")
	if err != nil {
		beego.Error("连接错误", err)
		beego.Info("连接错误", err)
		return
	}
	//关闭表
	defer conn.Close()

	//创建表
	_, err1 := conn.Exec("CREATE TABLE user (name varchar(40), password varchar(40));")
	if err1 != nil {
		beego.Error("创建表错误", err1)
		beego.Info("创建表错误", err1)
	}
}
