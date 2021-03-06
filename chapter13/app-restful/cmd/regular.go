// +build ignore

package main

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/sammyne/ppp-ddd/chapter13/app-restful/account.management/regular"
)

func accountService() *restful.WebService {
	service := new(restful.WebService)

	service.Path("/feeds")
	service.Route(service.GET("/").To(regular.Feeds))

	return service
}

func main() {
	restful.DefaultResponseContentType(restful.MIME_JSON)

	cors := restful.CrossOriginResourceSharing{Container: restful.DefaultContainer}
	restful.Filter(cors.Filter)

	restful.Add(accountService())

	server := &http.Server{
		Addr:    "localhost:4101",
		Handler: restful.DefaultContainer,
	}

	if err := server.ListenAndServe(); nil != err {
		panic(err)
	}
}
