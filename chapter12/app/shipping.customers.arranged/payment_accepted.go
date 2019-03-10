package arranged

import (
	"fmt"

	"github.com/sammyne/ppp-ddd/chapter12/app/billing.messages/events"
	shippingevents "github.com/sammyne/ppp-ddd/chapter12/app/shipping.messages/events"
)

func HandlePaymentAccepted(event *events.PaymentAccepted) {
	address := FindCustomerAddress(event.OrderID)

	cfrm := ArrangeShipping(address, event.OrderID)

	if cfrm.Status == Success {
		eventNew := &shippingevents.Arranged{OrderID: event.OrderID}

		brokerJSON.Publish("/shipping", eventNew)
		fmt.Printf("Shipping BC arranged shipping for Order: %s to: %s", event.OrderID, address)
		return
	}

	// TODO
}
