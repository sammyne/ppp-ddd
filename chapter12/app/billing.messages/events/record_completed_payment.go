package events

type PaymentStatus int

const (
	Accepted PaymentStatus = iota
	Rejected
)

type RecordPaymentAttempt struct {
	OrderID string
	Status  PaymentStatus
}
