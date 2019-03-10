package accepted

import (
	"github.com/sammyne/ppp-ddd/chapter12/app/billing.messages/events"
)

func FindCardDetails(userID string) *CardDetails {
	return new(CardDetails)
}

func SavePaymentAttempt(orderID string, status events.PaymentStatus) {

}
