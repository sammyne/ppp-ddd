package model

import (
	"errors"
	"time"
)

// Bid is the identity of a biding
type Bid struct {
	bidder     GUID
	maximumBid Money
	time       time.Time
}

func (bid *Bid) Bidder() GUID {
	return bid.bidder
}

func (bid *Bid) MaximumBid() Money {
	return bid.maximumBid
}

func (bid *Bid) TimeOfOffer() time.Time {
	return bid.time
}

func NewBid(bidderId GUID, maximumBid Money, time time.Time) (
	*Bid, error) {

	switch {
	case bidderId == EmptyGUID:
		return nil, errors.New("bidderId cannot be null")
	case maximumBid == NilMoney:
		return nil, errors.New("maximum bid cannot be null")
	case time.Before(minTime):
		return nil, errors.New("time of offer must have a value")
	}

	return &Bid{bidderId, maximumBid, time}, nil
}
