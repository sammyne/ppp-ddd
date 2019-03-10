package arranged

var orders []*ShippingOrder

func addOrder(order *ShippingOrder) {
	orders = append(orders, order)
}

func FindCustomerAddress(orderID string) string {
	/*
	* Implement this properly by storing the user's address
	* when the Sales bounded context publishes an event
	* with their details (add that message yourself, too)
	 */
	for _, o := range orders {
		if o.ID == orderID {
			return o.UserID + ", 55 DDDesign Street"
		}
	}

	return ""
}
