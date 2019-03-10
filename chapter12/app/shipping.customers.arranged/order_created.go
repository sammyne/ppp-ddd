package arranged

import (
	"fmt"

	"github.com/sammyne/ppp-ddd/chapter12/app/sales/messages/events"
)

func HandleOrderCreated(order *events.OrderCreated) {
	fmt.Printf("Shipping BC storing: Order: %s, User: %s, Shipping Type: %s\n", order.ID, order.UserID, order.ShippingTypeID)

	o := &ShippingOrder{
		UserID:         order.UserID,
		ID:             order.ID,
		ShippingTypeID: order.ShippingTypeID,
	}
	addOrder(o)
}
