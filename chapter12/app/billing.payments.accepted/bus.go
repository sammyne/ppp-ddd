package accepted

import nats "github.com/nats-io/go-nats"

var broker *nats.Conn

func init() {
	var err error
	broker, err = nats.Connect(nats.DefaultURL)
	if nil != err {
		panic(err)
	}

	broker.Subscribe("/", HandleOrderCreated)
	broker.Subscribe("/billing/local", HandleRecordPaymentAttempt)
}
