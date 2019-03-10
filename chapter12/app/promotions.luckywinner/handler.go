package promotions

import "fmt"

func HandleOrderCreated(msg *OrderCreated) {
	fmt.Printf("message-bus handling order placed event: Order: %s for User: %s", msg.ID, msg.UserID)
}
