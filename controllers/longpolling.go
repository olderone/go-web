package controllers

import(
	"web/models"
	"github.com/astaxie/beego"
	"time"
)

type LongPollingController struct{
	baseController
}

func init() {
	beego.Trace("开启了协程: ")
	go chatroom()
}

func (this *LongPollingController) Join() {
	// 安全检测
	uname := this.GetString("uname")
	if(len(uname) == 0){
		this.Redirect("/chartRoom",302)
		return
	}

	// 进入房间
	Join(uname, nil)

	this.TplName = "home/chart/lp.html"
	this.Data["IsLongPolling"] = true
	this.Data["UserName"] = uname
}

func (this *LongPollingController) Post() {
	this.TplName = "home/chart/lp.html"

	uname := this.GetString("uname")
	content := this.GetString("content")
	if len(uname) == 0 || len(content) == 0 {
		return
	}

	publish <- newEvent(models.EVENT_MESSAGE, uname, content, "")
}

//处理LongPollingController的获取归档请求.
func (this *LongPollingController) Fetch() {
	lastReceived, err := this.GetInt("lastReceived")
	if err != nil {
		return
	}

	events := models.GetEvents(int(lastReceived))
	if len(events) > 0 {
		for index := range events{
	      	events[index].Time = time.Unix(int64(events[index].Timestamp),0).Format("2006-01-02 15:04:05")
	   	}
		this.Data["json"] = events
		this.ServeJSON()
		return
	}

	// Wait for new message(s).
	ch := make(chan bool)
	waitingList.PushBack(ch)
	<-ch

	this.Data["json"] = models.GetEvents(int(lastReceived))
	this.ServeJSON()
}