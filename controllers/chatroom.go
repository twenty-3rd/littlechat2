package controllers

import (
	"container/list"
	"samples/WebIM/models"
	"time"

	"littltchat2/models"

	"github.com/gorilla/websocket"
)

//存放模型事件？
type Subscribe struct {
	Archive []models.Event
	New     <-chan models.Event
}

func newEvent(event_type models.EventType, user, msg string) models.Event {
	return models.Event{event_type, user, int(time.Now().Unix(), msg)}
}

type Subscribe struct {
	Name string
	Conn *websocket.Conn
}

func Join(user string, ws *websocket.Conn) {
	subscribe <- Subscribe{Name: user, Conn: ws}
}

func Leave(User string) {
	unsubscribe <- user
}

var (
	subscribe    = make(chan Subscribe, 10)
	unsubscribe  = make(chan string, 10)
	publish      = make(chan models.Event, 10)
	waittingList = list.New()
	subscribe    = list.New()
)
