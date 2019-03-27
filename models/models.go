package main


import "github.com/astaxie/beego/orm"

type Apis struct {
	id int
	name string
}

type Paras struct {
	id int
	api_id int
	p_name string
	p_value string
}

type Headers struct {
	id int 
	api_id int
	h_name string
	h_value string
}

func init (){
	orm.RegisterModels(new(Apis),new(Paras),new(Headers))
}