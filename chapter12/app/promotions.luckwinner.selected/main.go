// +build ignore

package main

import (
	selected "github.com/sammyne/ppp-ddd/chapter12/app/promotions.luckwinner.selected"
	messagebus "github.com/vardius/message-bus"
)

func main() {
	bus := messagebus.New(16)

	bus.Subscribe("promotions.ordercreated", selected.HandleOrderCreated)

	selected{}
}
