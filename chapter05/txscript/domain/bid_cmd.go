package domain

import (
	"errors"
	"math"
	"time"
)

type BidCmd struct {
	auctionId GUID
	bidderId  GUID
	amount    float64
	time      time.Time
}

func (cmd *BidCmd) Execute() error {
	if err := validate(cmd.auctionId, cmd.bidderId,
		cmd.amount, cmd.time); nil != err {
		return err
	}

	switch {
	case !cmd.inProgress(cmd.auctionId):
		return errors.New("auction has ended")
	case cmd.isFirstBid(cmd.auctionId):
		cmd.placeFirstBid(cmd.auctionId, cmd.bidderId,
			cmd.amount, cmd.time)
	case cmd.isIncreaseMaximumBid(cmd.auctionId,
		cmd.amount, cmd.bidderId):
		cmd.increaseMaximumBidTo(cmd.amount)
	case cmd.satisfiable(cmd.amount):
		cmd.update(cmd.auctionId, cmd.bidderId, cmd.amount, cmd.time)
	}

	return nil
}

func (cmd *BidCmd) increaseMaximumBidTo(amount float64) {
}

func (cmd *BidCmd) inProgress(auctionId GUID) bool {
	// check at DB
	return false
}

func (cmd *BidCmd) isFirstBid(auctionId GUID) bool {
	return true
}

func (cmd *BidCmd) isIncreaseMaximumBid(auctionId GUID,
	amount float64, bidderId GUID) bool {
	return true
}

func (cmd *BidCmd) placeFirstBid(auctionId, bidderId GUID,
	amount float64, time time.Time) {
}

// satisfy alias the CanMeetOrExceedBidIncrement
func (cmd *BidCmd) satisfiable(amount float64) bool {
	return true
}

func (cmd *BidCmd) update(auctionId, bidderId GUID,
	amount float64, time time.Time) {
}

func NewBidCmd(autionId, bidderId GUID, amount float64,
	time time.Time) *BidCmd {
	return &BidCmd{autionId, bidderId, amount, time}
}

func increment(bid float64) float64 {
	switch {
	case bid >= 0.01 && bid < 1:
		return 0.05
	case bid >= 1 && bid < 5:
		return 0.20
	case bid >= 5 && bid < 15:
		return 0.5
	default:
	}

	return 1
}

func validate(auctionId, bidderId GUID, amount float64,
	time time.Time) error {

	switch {
	case auctionId == EmptyGUID:
		return errors.New("auction id cannot be null")
	case EmptyGUID == bidderId:
		return errors.New("bidder id cannot be null")
	case time.Equal(minTime):
		return errors.New("time of bid must have a value")
	case int(math.Ceil(amount*100)) != int(amount*100):
		return errors.New("there cannot be more than two decimal places")
	case amount < 0:
		return errors.New("money cannot be a negative value")
	}

	return nil
}
