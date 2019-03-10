package promotions

import (
	"fmt"

	"github.com/sammyne/ppp-ddd/chapter12/app/sales/messages/events"
)

func BridgeOrderCreated(msg *events.OrderCreated) {
	fmt.Printf("Bridge received order: %s. About to push it onto message-bus's queue\n", msg.ID)

	event := &OrderCreated{
		ID:             msg.ID,
		UserID:         msg.UserID,
		ProductIDs:     msg.ProductIDs,
		ShippingTypeID: msg.ShippingTypeID,
		Timestamp:      msg.Timestamp,
		Amount:         msg.Amount,
	}

	mbus.Publish("promotions.ordercreated", event)
}
