package accepted

import (
	"encoding/json"
	"fmt"

	"github.com/sammyne/ppp-ddd/chapter12/app/bus"

	nats "github.com/nats-io/go-nats"
	"github.com/sammyne/ppp-ddd/chapter12/app/sales/messages/events"

	billingevents "github.com/sammyne/ppp-ddd/chapter12/app/billing.messages/events"
)

func HandleOrderCreated(msg *nats.Msg) {
	var order events.OrderCreated
	if err := json.Unmarshal(msg.Data, &order); nil != err {
		panic(err)
	}

	fmt.Println("Received order created event: OrderId:", order.ID)

	card := FindCardDetails(order.UserID)
	conf := ChargeCreditCard(card, order.Amount)
	cmd := &billingevents.RecordPaymentAttempt{
		OrderID: order.ID,
		Status:  conf.Status,
	}

	bus.PublishTo(broker, "billing/local", cmd)
}
