package controllers

import (
	"container/list"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"web/models"
)

type Subscription struct {
	Archive []models.Event      // 来自存档的所有事件.
	New     <-chan models.Event // 新事件接踵而至.
}

func newEvent(ep models.EventType, user, msg string, ts string) models.Event {
	return models.Event{ep, user, int(time.Now().Unix()), msg, ts}
}

func Join(user string, ws *websocket.Conn) {
	subscribe <- Subscriber{Name: user, Conn: ws}
}

func Leave(user string) {
	unsubscribe <- user
}

type Subscriber struct {
	Name string
	Conn *websocket.Conn // 只适用于WebSocket用户;否则零.
}

var (
	// 新加入用户的通道.
	subscribe = make(chan Subscriber, 10)
	// 退出的用户的通道.
	unsubscribe = make(chan string, 10)
	// 发送事件到这里发布它们.
	publish = make(chan models.Event, 10)
	// 长轮询等待列表.
	waitingList = list.New()
	subscribers = list.New()
)

// 此函数处理所有传入的chan消息.
func chatroom() {
	for {
		select {
		case sub := <-subscribe:
			if !isUserExist(subscribers, sub.Name) {
				subscribers.PushBack(sub) // 将用户添加到列表的末尾.
				// 发布连接事件.
				publish <- newEvent(models.EVENT_JOIN, sub.Name, "", "")
				beego.Info("New user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			} else {
				beego.Info("Old user:", sub.Name, ";WebSocket:", sub.Conn != nil)
			}
		case event := <-publish:
			// Notify waiting list.
			for ch := waitingList.Back(); ch != nil; ch = ch.Prev() {
				ch.Value.(chan bool) <- true
				waitingList.Remove(ch)
			}

			broadcastWebSocket(event)
			models.NewArchive(event)

			if event.Type == models.EVENT_MESSAGE {
				beego.Info("Message from", event.User, ";Content:", event.Content)
			}
		case unsub := <-unsubscribe:
			for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Name == unsub {
					subscribers.Remove(sub)
					// Clone connection.
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						beego.Error("WebSocket closed:", unsub)
					}
					publish <- newEvent(models.EVENT_LEAVE, unsub, "", "") // Publish a LEAVE event.
					break
				}
			}
		}
	}
}

func init() {
	beego.Trace("开启了协程: ")
	go chatroom()
}

// 判断用户是否还在排队
func isUserExist(subscribers *list.List, user string) bool {
	for sub := subscribers.Front(); sub != nil; sub = sub.Next() {
		if sub.Value.(Subscriber).Name == user {
			return true
		}
	}
	return false
}
