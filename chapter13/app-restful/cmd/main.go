package main

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	acontroller "github.com/sammyne/ppp-ddd/chapter13/app-restful/account.management/controllers"
	"github.com/sammyne/ppp-ddd/chapter13/app-restful/controllers"
)

func accountService() *restful.WebService {
	service := new(restful.WebService)

	service.Path("/accountmanagement/accounts")

	service.GET("/").To(acontroller.Index)
	service.GET("/{accountID}").To(acontroller.Account)

	return service
}

func entryPointService() *restful.WebService {
	service := new(restful.WebService)

	service.Path("/accountmanagement").Route(service.GET("/").To(controllers.GetEntryPoint))

	return service
}

func main() {
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.Add(entryPointService())

	server := &http.Server{
		Addr:    "localhost:4100",
		Handler: restful.DefaultContainer,
	}

	if err := server.ListenAndServe(); nil != err {
		panic(err)
	}
}
