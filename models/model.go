package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//表的设计

//定义一个user的结构体
type User struct {
	Id       int
	Name     string
	Password string
}

//文章表
type Article struct {
	Id       int       `orm:"pk;auto"`
	ArtiName string    `orm:"size(20)"`
	Atime    time.Time `orm:"auto_now"`
	Acount   int       `orm:"default(0);null"`
	Acontent string    `orm:"size(500)"`
	Aimg     string    `orm:"size(100)"`
}

//写init这个函数名是为了main函数引入包的时候就可以自动调用
func init() {
	//orm操作数据库
	orm.RegisterDataBase("default", "mysql", "root:@tcp(10.10.10.201:3306)/test?charset=utf8")
	//创建表
	orm.RegisterModel(new(User), new(Article))

	//生成表,第一个参数：数据库别名，第二个：是否强制更新，第三个：是否可见过程
	orm.RunSyncdb("default", false, true)

	//操作数据库代码
	/*conn, err := sql.Open("mysql", "root:@tcp(10.10.10.201:3306)/test?charset=utf8")
	if err != nil {
		beego.Error("连接错误", err)
		beego.Info("连接错误", err)
		return
	}
	//关闭表
	defer conn.Close()

	//创建表
	/*_, err1 := conn.Exec("CREATE TABLE user (name varchar(40), password varchar(40));")*/
	//插入数据
	//_, err1 := conn.Exec("insert into user(name, password) values (?, ?)", "chuanzhi", "heima")
	/*res, err1 := conn.Query("select name from user")
	var name string
	for res.Next() {
		res.Scan(&name)
		beego.Info(name)
	}
	if err1 != nil {
		beego.Error("创建表错误", err1)
		beego.Info("创建表错误", err1)
	}*/
}
