package accepted

import (
	"github.com/sammyne/ppp-ddd/chapter12/app/billing.messages/events"
)

func ChargeCreditCard(details *CardDetails, amount float64) *PaymentConfirmation {
	return &PaymentConfirmation{events.Accepted}
}
