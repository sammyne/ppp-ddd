package main

import (
	"net/http"

	"github.com/sammyne/ppp-ddd/chapter12/app/web/controllers"
)

func main() {
	http.HandleFunc("/orders", controllers.OrdersIndex)
	http.HandleFunc("/orders/place", controllers.PlaceOrder)

	if err := http.ListenAndServe(":8080", nil); nil != err {
		panic(err)
	}
}
