package domain

import "time"

type Auction struct {
	Id          GUID
	ListingId   GUID
	EndAt       time.Time
	Base        float64
	WinnerPrice float64
	MaximumBid  float64
	Winner      GUID
}
