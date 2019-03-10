package created

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/sammyne/ppp-ddd/chapter12/app/bus"

	nats "github.com/nats-io/go-nats"
	"github.com/sammyne/ppp-ddd/chapter12/app/sales/messages/commands"
	"github.com/sammyne/ppp-ddd/chapter12/app/sales/messages/events"
)

func HandlePlaceOrder(msg *nats.Msg) {
	var cmd commands.PlaceOrder
	if err := json.Unmarshal(msg.Data, &cmd); nil != err {
		panic(err)
	}

	orderID := saveOrder(cmd.ProductIDs, cmd.UserID, cmd.ShippingTypeID)
	fmt.Printf("Created order %s for Products: %s with shipping: %s made by user: %s\n",
		orderID, strings.Join(cmd.ProductIDs, ","), cmd.ShippingTypeID, cmd.UserID)

	event := &events.OrderCreated{
		ID:             orderID,
		UserID:         cmd.UserID,
		ProductIDs:     cmd.ProductIDs,
		ShippingTypeID: cmd.ShippingTypeID,
		Timestamp:      time.Now(),
		Amount:         calculateCost(cmd.ProductIDs),
		AddressID:      "AddressID123",
	}

	bus.PublishTo(broker, "/", event)
}

func calculateCost(productIDs []string) float64 {
	return 1000
}
