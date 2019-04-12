package models

import (
	"github.com/astaxie/beego/orm"
)

type CityInfo struct {
	Id           int32 `orm:"auto"`
	CityNo       int32
	CityName     string
	CountryName  string
	ProvinceName string
}

func GetCityList() (city []CityInfo, err error, nu int64) {
	//qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	//	qb.Select("city.cityName",
	//		"province.provinceName").
	//		From("city").
	//		LeftJoin("province").On("city.procinceId = province.provinceNo")

	// 导出 SQL 语句
	//sql := qb.String()

	o := orm.NewOrm()
	o.Using("ec")
	//o.Raw(sql).QueryRows(&city)
	num, err := o.Raw("SELECT city.id,city.CityNo,CityName,countryName,provinceName FROM city left join province AS pro ON city.provinceId = pro.provinceNo left join country AS ctry ON ctry.countryNo = city.countryId").QueryRows(&city)

	return city, err, num
}
