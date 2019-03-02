package controllers

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/nvellon/hal"
	"github.com/sammyne/ppp-ddd/chapter13/app-restful/conf"
)

const (
	entryPointBaseURL = "http://localhost:4100/"
	accountsBaseURL   = "http://localhost:4101/accountmanagement/"
)

func Index(req *restful.Request, resp *restful.Response) {
	root := hal.NewResource(struct{}{}, accountsBaseURL)
	root.AddLink("alternative", hal.NewLink(accountsBaseURL+"accounts?page=1"))
	root.AddLink("account", hal.NewLink(accountsBaseURL+"accounts/{accountID}"))
	root.AddLink("account", hal.NewLink(accountsBaseURL+"accounts/123"))
	root.AddLink("account", hal.NewLink(accountsBaseURL+"accounts/456"))
	root.AddLink("next", hal.NewLink(accountsBaseURL+"accounts?page=2"))
	root.AddLink("parent", hal.NewLink(entryPointBaseURL))

	resp.WriteHeaderAndJson(http.StatusOK, root, "application/hal+json")
}

func Account(req *restful.Request, resp *restful.Response) {
	resp.Header().Set("Content-Type", conf.MIME_HAL_JSON)

	/*
		root := hal.NewResource(struct{}{}, entryPointBaseURL+"/accountmanagement")
		root.AddLink("accounts", hal.NewLink(accountsBaseURL+"/accountmanagement/accounts"))

		data, _ := json.Marshal(root)
		resp.Write(data)
	*/
}
