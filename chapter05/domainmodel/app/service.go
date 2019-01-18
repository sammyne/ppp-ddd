package app

import "github.com/sammyne/ppp-ddd/chapter05/domainmodel/model"
import "time"

type BidOnAuctionService struct {
	repository model.AuctionRepository
}

func (service *BidOnAuctionService) Bid(autionId, memberId model.GUID,
	amount float64, time time.Time) error {
	auction, err := service.repository.Find(autionId)
	if nil != err {
		return err
	}

	offer, err := model.NewBid(memberId, model.Money(amount), time)
	if nil != err {
		return err
	}

	return auction.PlaceBidFor(*offer, time)
}

func NewBidOnActionService(
	repository model.AuctionRepository) *BidOnAuctionService {
	return &BidOnAuctionService{repository}
}
