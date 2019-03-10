package events

type PaymentStatus int

const (
	Accepted PaymentStatus = iota
	Rejected
)

type RecordCompletedPayment struct {
	OrderID string
	Status  PaymentStatus
}
