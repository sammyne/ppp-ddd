package commands

import "time"

type PlaceOrder struct {
	userID         string
	productIDs     []string
	shippingTypeID string
	timestamp      time.Time
}

func (cmd *PlaceOrder) GetUserID() string {
	return cmd.userID
}

func (cmd *PlaceOrder) GetProductIDs() []string {
	return cmd.productIDs
}

func (cmd *PlaceOrder) GetShippingTypeID() string {
	return cmd.shippingTypeID
}

func (cmd *PlaceOrder) GetTimestamp() time.Time {
	return cmd.timestamp
}

func (cmd *PlaceOrder) SetUserID(id string) {
	cmd.userID = id
}

func (cmd *PlaceOrder) SetProductIDs(ids []string) {
	cmd.productIDs = ids
}

func (cmd *PlaceOrder) SetShippingTypeID(id string) {
	cmd.shippingTypeID = id
}

func (cmd *PlaceOrder) SetTimestamp(timestamp time.Time) {
	cmd.timestamp = timestamp
}
