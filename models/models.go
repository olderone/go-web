package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func RegistDB() {
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	//注册 model
	orm.RegisterModel(new(HospInfo), new(AdminUsers))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:root@/ec_hosp?charset=utf8") //密码为空格式

}
