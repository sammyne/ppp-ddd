package model

import (
	"errors"
	"time"
)

// Auction is an Auction Class from a Rich Domain Model
type Auction struct {
	base      Money // starting price
	bids      []*HistoricalBid
	endAt     time.Time
	ended     bool
	id        GUID
	listingId GUID
	winner    *WinningBid // WinningBid
}

func (aution *Auction) CanPlaceBid() bool {
	return !aution.ended
}

func (aution *Auction) PlaceBidFor(bid Bid, now time.Time) error {
	var err error
	switch {
	case !aution.inProgress(now):
	case nil == aution.winner: // 1st bid
		aution.registerFirst(bid)
	case aution.better(bid):
		aution.winner, err = aution.winner.RaiseMaximumBidTo(bid.maximumBid)
	case aution.winner.price.Increment() <= bid.maximumBid:
		aution.place(aution.winner.DetermineWinningBidIncrement(bid))
	default:
	}

	return err
}

// inProgress checks if the aution is still active, i.e.,
// its ending time is before the provided now
func (aution *Auction) inProgress(now time.Time) bool {
	return aution.endAt.After(now)
}

// better checks if the given bid is better than the one up till now
func (aution *Auction) better(bid Bid) bool {
	var ok bool
	switch {
	case nil == aution.winner:
	case !aution.winner.MadeBy(bid.bidder):
	case bid.maximumBid <= aution.winner.maximumBid:
	default:
		ok = true
	}

	return ok
}

// Place adds the next winner bid
func (aution *Auction) place(bid *WinningBid) {
	aution.bids = append(aution.bids,
		NewHistoricalBid(bid.bidder, bid.maximumBid, bid.time))
	aution.winner = bid
}

func (aution *Auction) registerFirst(bid Bid) {
	if nil != aution.winner {
		return // not the 1st
	}

	if bid.maximumBid < aution.base {
		return
	}

	winner, _ := NewWinningBid(bid.bidder, bid.maximumBid,
		aution.base, bid.time)
	aution.place(winner)
}

func NewAuction(id GUID, listingId GUID, base Money,
	endAt time.Time) (*Auction, error) {

	switch {
	case EmptyGUID == id:
		return nil, errors.New("auction id cannot be null")
	case NilMoney == base:
		return nil, errors.New("starting price cannot be null")
	case minTime == endAt:
		return nil, errors.New("endAt must have a value")
	case EmptyGUID == listingId:
		return nil, errors.New("listing id cannot be null")
	default:
	}

	return &Auction{
		base:      base,
		endAt:     endAt,
		id:        id,
		listingId: listingId,
	}, nil
}
