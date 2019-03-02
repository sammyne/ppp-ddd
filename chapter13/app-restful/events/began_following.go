package events

type BeganFollowing struct {
	AccountID  string
	FollowerID string
}

func NewBeganFollowing(account, follower string) *BeganFollowing {
	return &BeganFollowing{AccountID: account, FollowerID: follower}
}
