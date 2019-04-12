package models

import (
	"github.com/astaxie/beego/orm"
)

type ExtraInfo struct {
	Fid      string
	MapNo    int32
	GroupId  int32
	Type     int32
	Floor    string
	Room     string
	OrgZoom  string
	OrgNo    int32
	OrgName  string
	ParkZoom string
	Height   float32
	Name     string
	X        float32
	Y        float32
	Area     float32
	Url      string
}

func GetExtra(hospNo int) (extra []ExtraInfo, err error, nu int64) {
	o := orm.NewOrm()
	num, err := o.Raw("SELECT hml.fid,hml.map_no,hml.group_id,hml.type, hml.floor, hml.room, hml.org_zone, hml.org_no, oh.name as OrgName, hml.park_zone, hml.height, hml.name, hml.x, hml.y, hml.area, hml.name as descp, hior.url FROM hosp_map_layer AS hml right join hosp_exit_img AS hei ON hml.fid = hei.fid left join hosp_img_own_res AS hior ON hei.img_no = hior.img_no left join org_hosp AS oh ON hml.org_no=oh.org_no WHERE hml.hosp_no = ?", hospNo).QueryRows(&extra)

	return extra, err, num
}
