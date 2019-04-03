package main

import (
	"time"
	_ "web/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//引入数据模型
func init() {

	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/ec_hosp?charset=utf8") //密码为空格式
}

func main() {
	//这时, 在你重启应用的时候, beego 便会自动帮你创建数据库表。
	orm.Debug = true

	orm.RunSyncdb("default", false, true)

	beego.Run()
}
