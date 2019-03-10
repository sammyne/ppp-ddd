package accepted

import (
	"encoding/json"
	"fmt"

	nats "github.com/nats-io/go-nats"
	"github.com/sammyne/ppp-ddd/chapter12/app/sales/messages/events"
)

func HandleOrderCreated(msg *nats.Msg) {
	var order events.OrderCreated
	if err := json.Unmarshal(msg.Data, &order); nil != err {
		panic(err)
	}

	fmt.Println("Received order created event: OrderId:", order.ID)
}
