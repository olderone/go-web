package models

import(
	"fmt"
	"container/list"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct{
	Type 		EventType
	User 	 	string
	Timestamp 	int
	Content 	string
	Time 		string
}

type ChartRecord struct {
	Id         int32  `orm:"auto"`
	Type     EventType  //园区类型
	Timestamp     int  //城市编号
	User       string `orm:"size(100)"`  //性别
	Content      string `orm:"size(255)"` //Email
	Contact    string `orm:"size(255)"` //Email
	Time        string `orm:"size(255)"` //电话
}


const archiveSize = 30

//事件驱动
var archive = list.New()

func typeof(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

// NewArchive将新事件保存到存档列表.
func NewArchive(event Event) {
	o := orm.NewOrm()
	var cr ChartRecord
	cr.Type = event.Type
	cr.Timestamp = event.Timestamp
	cr.User = event.User
	cr.Content = event.Content
	cr.Time = event.Time
	if(cr.Content != ""){
		o.Insert(&cr)
	}
	
	beego.Trace("开始1")
	beego.Trace(event)
	beego.Trace(event.Type)
	beego.Trace(typeof(event))
	beego.Trace("结束1")
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}

// GetEvents返回lastReceived之后的所有事件.
func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	beego.Trace("开始")
	beego.Trace(events)
	beego.Trace(typeof(events))
	beego.Trace("结束")
	return events
}