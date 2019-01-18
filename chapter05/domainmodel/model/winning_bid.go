package model

import (
	"errors"
	"time"
)

type WinningBid struct {
	actionId   GUID
	bidder     GUID
	maximumBid Money
	time       time.Time // TimeOfBid
	price      Price     // CurrentAuctionPrice
}

//func (bid *WinningBid)

func (bid *WinningBid) AutionID() GUID {
	return bid.actionId
}

func (bid *WinningBid) Bidder() GUID {
	return bid.bidder
}

//
func (bid *WinningBid) DetermineWinningBidIncrement(other Bid) *WinningBid {

	if bidMax := bid.price.Increment(); bid.maximumBid >= bidMax && other.maximumBid >= bidMax {

		if bid.maximumBid < other.maximumBid {

			next := Price(bid.maximumBid).Increment()
			if next >= other.maximumBid {
				next = other.maximumBid
			}

			winner, _ := NewWinningBid(other.bidder, next, other.maximumBid, other.time)

			return winner
		}

		next := Price(other.maximumBid).Increment()
		if next >= bid.maximumBid {
			next = bid.maximumBid
		}

		winner, _ := NewWinningBid(bid.bidder, next, bid.maximumBid, bid.time)

		return winner
	} else if other.maximumBid >= bidMax {
		winner, _ := NewWinningBid(other.bidder, bidMax,
			other.maximumBid, other.time)

		return winner
	}

	return bid
}

func (bid *WinningBid) MaximumBid() Money {
	return bid.maximumBid
}

func (bid *WinningBid) Price() Price {
	return bid.price
}

func (bid *WinningBid) RaiseMaximumBidTo(amount Money) (
	*WinningBid, error) {
	if amount <= bid.maximumBid {
		return nil, errors.New("maximum bid increase must be larger than current maximum bid")
	}

	return NewWinningBid(bid.bidder, amount, Money(bid.price), time.Now())
}

func (bid *WinningBid) Time() time.Time {
	return bid.time
}

func (bid *WinningBid) MadeBy(bidder GUID) bool {
	return bidder == bid.bidder
}

func CopyTime(timestamp time.Time) time.Time {
	data, _ := timestamp.MarshalBinary()

	timestampCopy := new(time.Time)
	timestampCopy.UnmarshalBinary(data)

	return *timestampCopy
}

func NewWinningBid(bidder GUID, maximumBid Money, price Money,
	time time.Time) (*WinningBid, error) {
	switch {
	case bidder == EmptyGUID:
		return nil, errors.New("bidder cannot be null")
	case NilMoney == maximumBid:
		return nil, errors.New("maximum bid cannot be null")
	case minTime.Equal(time):
		return nil, errors.New("time of bidding must have a value")
	}

	return &WinningBid{
		bidder:     bidder,
		maximumBid: maximumBid,
		time:       time,
		price:      Price(price),
	}, nil
}

//func newWinningBid(bidder GUID, bid, maxBid Money, time time.Time) (
//	*WinningBid,error) {
//	return N
//}
