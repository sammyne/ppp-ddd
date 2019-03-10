package accepted

import (
	"github.com/sammyne/ppp-ddd/chapter12/app/billing.messages/events"
)

type PaymentConfirmation struct {
	Status events.PaymentStatus
}
