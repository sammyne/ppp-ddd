package promotions

import "time"

type OrderCreated struct {
	ID             string
	UserID         string
	ProductIDs     []string
	ShippingTypeID string
	Timestamp      time.Time
	Amount         float64
}
