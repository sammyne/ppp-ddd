package model

import "time"

type HistoricalBid struct {
	Bidder GUID
	Amount Money
	Time   time.Time
}

func NewHistoricalBid(bidder GUID, amount Money,
	time time.Time) *HistoricalBid {
	return &HistoricalBid{bidder, amount, time}
}
