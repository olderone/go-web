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
	ImgUrl     string `orm:"size(255)"` //电话
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

type InsPak struct {
	Id int32 `orm:"auto"`
	//VersionName string `orm:"size(32)"`  //版本名称
	//UpdateCont  string `orm:"size(255)"` //更新内容
	Url string `orm:"size(255)"` //下载地址
	//UrlQiniu    string `orm:"size(255)"` //七牛云下载地址
	//PakSize     int32  //包大小
	//IsForce     int8   //是否强制更新
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

// 安装包
func GetInstallPak(appver string) (ip InsPak, err error) {
	o := orm.NewOrm()

	//num, err := o.Raw("SELECT id,url FROM hosp_version WHERE apk_name > ? ORDER BY created_at DESC LIMIT 1, 2", '1.0').QueryRow(&ip)
	//var users []User
	//res, err := o.Raw("SELECT  id, url FROM hosp_version WHERE id > ?", appver).Exec()
	error := o.Raw("SELECT  id, url FROM hosp_version WHERE version_name > ?", appver).QueryRow(&ip)

	return ip, error

}

//func query() {
//	db, err := sql.Open("mysql", "root:root@/golang?charset=utf8")
//	checkErr(err)

//	rows, err := db.Query("SELECT * FROM user")
//	checkErr(err)

//	//    //普通demo
//	for rows.Next() {
//		var userid int
//		var username string
//		var userage int
//		var usersex int

//		rows.Columns()
//		err = rows.Scan(&userid, &username, &userage, &usersex)
//		checkErr(err)

//		fmt.Println(userid)
//		fmt.Println(username)
//		fmt.Println(userage)
//		fmt.Println(usersex)
//	}
//}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
