package models

import (
	"container/list"
	"fmt"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct {
	Type      EventType
	User      string
	Timestamp int
	Content   string
}

const archiveSize = 20

var archive = list.New()

func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}
	archive.PushBack(event)
}

func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		fmt.Printf("=============%v", e)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
