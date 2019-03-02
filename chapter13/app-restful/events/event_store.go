package events

import (
	"bytes"
	"encoding/json"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const eventStoreURL = "http://localhost:2113"

type Event struct {
	ID   string          `json:"eventId"`
	Type string          `json:"eventType"`
	Data json.RawMessage `json:"data"`
}

//type EventPersister struct {
//}

//func (persister *EventPersister) Persist(data interface{}) {
func Persist(data interface{}) {
	eventJSON, _ := json.Marshal(NewEvent(data, "BeganFollowing"))

	req, err := http.NewRequest("POST", eventStoreURL+"/streams/BeganFollowing", bytes.NewReader(eventJSON))
	if nil != err {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ES-CurrentVersion", "1")

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		panic(err)
	}
	defer resp.Body.Close()
}

func NewEvent(data interface{}, eventType string) *Event {
	dataJSON, _ := json.Marshal(data)
	return &Event{
		ID:   uuid.Must(uuid.NewV4()).String(),
		Type: eventType,
		Data: dataJSON,
	}
}
