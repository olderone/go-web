package controllers

import(
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

var langTypes []string //支持的语言

// 初始化加载配置文件
func init() {
	// 初始化语言类型列表
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	// 根据预言类型加载locale文件
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

// baseController表示所有其他应用程序路由器的基本路由器.
// 它为相同的实现实现了一些方法;
// 因此，它将被嵌入到其他路由器.
type baseController struct {
	beego.Controller // 具有接口存根实现的嵌入结构.
	i18n.Locale      // 用于处理数据和呈现模板时使用i18n.
}

// 为baseController实现了Prepare()方法.
// 用于语言选项的检查和设置.
func (this *baseController) Prepare() {
	// 重新设置语言选项.
	this.Lang = "" // 这个字段来自i18n.Locale.

	// 1. 从“Accept-Language”中获取语言信息'.
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // 只比较前5个字母.
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}

	// 2. 默认语言English
	if len(this.Lang) == 0 {
		this.Lang = "zh-CN"
	}

	// 设置模板级语言选项.
	this.Data["Lang"] = this.Lang
}


type ChartController struct{
	baseController // 嵌入使用在baseController中实现的方法.
}


func (this *ChartController) Get() {
	this.TplName = "home/chart/welcome.html"
}

// Join方法处理AppController的POST请求.
func (this *ChartController) Join() {
	// Get form value.
	uname := this.GetString("uname")
	tech := this.GetString("tech")

	// Check valid.
	if len(uname) == 0 {
		this.Redirect("/chartRoom", 302)
		return
	}

	switch tech {
	case "longpolling":
		this.Redirect("/lp?uname="+uname, 302)
	case "websocket":
		this.Redirect("/ws?uname="+uname, 302)
	default:
		this.Redirect("/", 302)
	}

	// Usually put return after redirect.
	return
}
