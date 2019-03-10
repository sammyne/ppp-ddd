package created

import (
	"encoding/json"
	"fmt"
	"strings"

	nats "github.com/nats-io/go-nats"
	"github.com/sammyne/ppp-ddd/chapter12/app/sales/messages/commands"
)

func HandlePlaceOrder(msg *nats.Msg) {
	var cmd commands.PlaceOrder
	if err := json.Unmarshal(msg.Data, &cmd); nil != err {
		panic(err)
	}

	fmt.Printf("Order for Products: %s with shipping: %s made by user: %s\n",
		strings.Join(cmd.ProductIDs, ","), cmd.ShippingTypeID, cmd.UserID)
}
