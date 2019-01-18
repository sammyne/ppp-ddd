package model

type AuctionRepository interface {
	Add(item *Auction)
	Find(id GUID) (*Auction, error)
}
