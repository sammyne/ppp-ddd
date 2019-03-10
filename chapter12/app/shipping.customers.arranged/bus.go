package arranged

import nats "github.com/nats-io/go-nats"

var brokerJSON *nats.EncodedConn

func init() {
	broker, err := nats.Connect(nats.DefaultURL)
	if nil != err {
		panic(err)
	}

	brokerJSON, err = nats.NewEncodedConn(broker, "json")
	if nil != err {
		panic(err)
	}

	brokerJSON.Subscribe("Sales.Orders.OrderCreated", HandleOrderCreated)
	brokerJSON.Subscribe("Billing.Payments.PaymentAccepted", HandlePaymentAccepted)
}
