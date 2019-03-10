package promotions

import (
	nats "github.com/nats-io/go-nats"
	messagebus "github.com/vardius/message-bus"
)

var mbus messagebus.MessageBus
var natBus *nats.EncodedConn

func init() {
	mbus = messagebus.New(16)
	mbus.Subscribe("promotions.ordercreated", HandleOrderCreated)

	if broker, err := nats.Connect(nats.DefaultURL); nil != err {
		panic(err)
	} else {
		natBus, err = nats.NewEncodedConn(broker, "json")
		if nil != err {
			panic(err)
		}
	}

	natBus.Subscribe("/", BridgeOrderCreated)
}
