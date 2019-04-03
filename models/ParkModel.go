package models

import (
	"github.com/astaxie/beego/orm"
)

type HospInfo struct {
	Id         int32  `orm:"auto"`
	TypeNo     int32  //园区类型
	CityNo     int32  //城市编号
	CountyNo   int32  //用户名
	HospNo     int32  //生日
	Name       string `orm:"size(32)"`  //性别
	Intro      string `orm:"size(500)"` //Email
	Url        string `orm:"size(255)"` //电话
	MapName    string `orm:"size(255)"` //性别
	FmapUrl    string `orm:"size(500)"` //Email
	MapUrl     string `orm:"size(255)"` //电话
	MapSize    int64  //状态
	Class      string `orm:"size(255)"` //生日
	Addr       string `orm:"size(255)"` //性别
	Contact    string `orm:"size(255)"` //Email
	Lon        string `orm:"size(255)"` //电话
	Lat        string `orm:"size(255)"` //性别
	UpdatedNum int32  //状态
	CreatedAt  int32  //创建时间
	UpdatedAt  int32  //更新时间
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

// 查询单条数据
func GetParkById(id int32) (v *HospInfo, err error) {
	o := orm.NewOrm()
	v = &HospInfo{Id: id}
	if err = o.Read(v, "Id"); err == nil {
		return v, nil
	}
	return nil, err
}
