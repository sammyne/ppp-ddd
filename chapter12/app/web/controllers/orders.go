package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/sammyne/ppp-ddd/chapter12/app/bus"
	"github.com/sammyne/ppp-ddd/chapter12/app/sales/messages/commands"
)

func OrdersIndex(resp http.ResponseWriter, req *http.Request) {
	// paths are w.r.t. to the app root
	filename := "views/index.html"
	tmpl, err := template.ParseFiles(filename)
	if nil != err {
		panic(err)
	}

	//resp.Write
	//fmt.Fprintf(resp, "Hello World")
	tmpl.Execute(resp, nil)
}

func PlaceOrder(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	userID, shippingTypeID := req.PostFormValue("userID"), req.PostFormValue("shippingTypeID")

	productIDs := strings.Split(req.PostFormValue("productIDs"), ",")

	cmd := commands.NewPlaceOrder(userID, productIDs, shippingTypeID)
	bus.PublishTo(broker, "Sales.Orders.OrderCreated", cmd)

	fmt.Fprintf(resp, "Your order has been placed. You will receive email confirmation shortly.")
}
