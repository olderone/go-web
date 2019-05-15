package models

import(
	"container/list"
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

const archiveSize = 30

//事件驱动
var archive = list.New()

// NewArchive将新事件保存到存档列表.
func NewArchive(event Event) {
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
	return events
}