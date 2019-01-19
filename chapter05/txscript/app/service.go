package app

import (
	"errors"
	"time"

	"github.com/sammyne/ppp-ddd/chapter05/txscript/domain"
)

type BidOnAuctionService struct {
	commander *domain.BidCmd
}

func (service *BidOnAuctionService) Bid(auctionId, memberId domain.GUID, amount float64, time time.Time) error {
	//return service.commander.Execute(auctionId, memberId, amount, time)
	return errors.New("not implemented")
}

func NewBidOnAuctionService(
	commander *domain.BidCmd) *BidOnAuctionService {
	return &BidOnAuctionService{commander}
}
