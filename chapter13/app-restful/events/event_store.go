package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const eventStoreURL = "http://localhost:2113"

type Event struct {
	ID   string          `json:"eventId"`
	Type string          `json:"eventType"`
	Data json.RawMessage `json:"data"`
}

func Persist(data interface{}) {
	dataJSON, err := json.Marshal(data)
	if nil != err {
		panic(err)
	}

	req, err := http.NewRequest("POST", eventStoreURL+"/streams/BeganFollowing", bytes.NewReader(dataJSON))
	if nil != err {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ES-EventId", uuid.Must(uuid.NewV4()).String())
	req.Header.Set("ES-EventType", "BeganFollowing")
	req.Header.Set("ES-CurrentVersion", "1")

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		panic(err)
	}
	defer resp.Body.Close()

	/*
		out, err := ioutil.ReadAll(resp.Body)
		if nil != err {
			panic(err)
		}
	*/
}

func Retrieve(v interface{}) {
	req, err := http.NewRequest("GET", eventStoreURL+"/streams/BeganFollowing", nil)
	if nil != err {
		panic(err)
	}
	req.Header.Set("Accept", "application/vnd.eventstore.atom+json")

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		panic(err)
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}
	fmt.Printf("%s\n", out)
}

func NewEvent(data interface{}, eventType string) *Event {
	dataJSON, _ := json.Marshal(data)
	return &Event{
		ID:   uuid.Must(uuid.NewV4()).String(),
		Type: eventType,
		Data: dataJSON,
	}
}
