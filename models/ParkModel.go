package models

import (
	"github.com/astaxie/beego/orm"
)

type HospInfo struct {
	Id          int32  `orm:"auto"`
	typeNo      int32  `orm:"auto"`      //园区类型
	cityNo      int32  `orm:"auto"`      //城市编号
	countyNo    int32  `orm:"auto"`      //用户名
	hospNo      int32  `orm:"auto"`      //生日
	NAME        string `orm:"size(32)"`  //性别
	intro       string `orm:"size(500)"` //Email
	url         string `orm:"size(255)"` //电话
	mapName     string `orm:"size(255)"` //性别
	fmapUrl     string `orm:"size(500)"` //Email
	mapUrl      string `orm:"size(255)"` //电话
	mapSize     int64  `orm:"auto"`      //状态
	class       string `orm:"size(255)"` //生日
	addr        string `orm:"size(255)"` //性别
	contact     string `orm:"size(255)"` //Email
	lon         string `orm:"size(255)"` //电话
	lat         string `orm:"size(255)"` //性别
	updatedNum  int32  //状态
	createdTime int32  //创建时间
	updatedTime int32  //更新时间
}

//查询数据
func QueryAllParkInfo() (dataList []interface{}, err error) {
	var list []HospInfo
	o := orm.NewOrm()
	qs := o.QueryTable(new(HospInfo))

	//查询数据
	if _, err = qs.All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}
