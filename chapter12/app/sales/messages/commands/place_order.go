package commands

import "time"

type PlaceOrder struct {
	UserID         string
	ProductIDs     []string
	ShippingTypeID string
	Timestamp      time.Time
}

func NewPlaceOrder(userID string, productIDs []string, shippingTypeID string) *PlaceOrder {
	return &PlaceOrder{
		userID,
		productIDs,
		shippingTypeID,
		time.Now(),
	}
}
