package main

import (
	"web/models"
	_ "web/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//引入数据模型
func init() {
	// 注册数据库
	models.RegistDB()
}

func main() {
	//这时, 在你重启应用的时候, beego 便会自动帮你创建数据库表。
	orm.Debug = true

	orm.RunSyncdb("default", false, true)

	beego.Run()
}
