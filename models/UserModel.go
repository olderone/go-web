package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminUsers struct {
	Id             int32  `orm:"auto"`
	Userame        string `orm:"size(255)"` //登录名
	Password       string `orm:"size(255)"` //密码
	Name           string `orm:"size(255)"` //用户名
	avatar         string `orm:"size(255)"` //生日
	remember_token int8   //性别
	created_at     int32  //创建时间
	updated_at     int32  //更新时间
}

//查询数据
func QueryAllUserInfo() (dataList []interface{}, err error) {
	var list []AdminUsers
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	qs := o.QueryTable(new(AdminUsers))

	//查询数据
	if _, err = qs.All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}
