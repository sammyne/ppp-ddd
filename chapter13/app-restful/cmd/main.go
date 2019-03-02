package main

import (
	"net/http"

	"github.com/sammyne/ppp-ddd/chapter13/app/account.management/controllers"
	rcontrollers "github.com/sammyne/ppp-ddd/chapter13/app/discovery/controllers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/accountmanagement", nil)
	mux.HandleFunc("/api/followerdirectory/getusersfollowers", controllers.GetUsersFollowers)
	mux.HandleFunc("/api/recommender/getrecommendedusers", rcontrollers.Users)

	server := &http.Server{
		Addr:    "localhost:3200",
		Handler: mux,
	}

	if err := server.ListenAndServe(); nil != err {
		panic(err)
	}
}
