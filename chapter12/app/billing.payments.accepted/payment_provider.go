package accepted

import (
	"errors"

	"github.com/sammyne/ppp-ddd/chapter12/app/billing.messages/events"
)

var nAttempts int

func ChargeCreditCard(details *CardDetails, amount float64) (*PaymentConfirmation, error) {
	if nAttempts < 2 {
		nAttempts++
		return nil, errors.New("Service unavailable. Down for maintenance.")
	}

	return &PaymentConfirmation{events.Accepted}, nil
}
