package main

import (
	"net/http"

	"github.com/sammyne/ppp-ddd/chapter13/app/account.management/controllers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/followerdirectory/getusersfollowers", controllers.GetUsersFollowers)

	server := &http.Server{
		Addr:    "localhost:3200",
		Handler: mux,
	}

	if err := server.ListenAndServe(); nil != err {
		panic(err)
	}
}
