package accepted

import (
	"encoding/json"
	"fmt"

	"github.com/sammyne/ppp-ddd/chapter12/app/bus"

	nats "github.com/nats-io/go-nats"
	"github.com/sammyne/ppp-ddd/chapter12/app/billing.messages/events"
)

func HandleRecordPaymentAttempt(msg *nats.Msg) {
	var attempt events.RecordPaymentAttempt
	if err := json.Unmarshal(msg.Data, &attempt); nil != err {
		panic(err)
	}

	SavePaymentAttempt(attempt.OrderID, attempt.Status)

	if events.Accepted == attempt.Status {
		event := events.PaymentAccepted{
			OrderID: attempt.OrderID,
		}
		bus.PublishTo(broker, "/payments", event)
		fmt.Printf("Received payment accepted notification for Order: %s. Published PaymentAccepted event\n", attempt.OrderID)

		return
	}

	// publish a payment rejected event
}
