package controllers

import (
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/nvellon/hal"
	"github.com/sammyne/ppp-ddd/chapter13/app-restful/conf"
)

const (
	entryPointBaseURL         = "http://localhost:4100/"
	accountsManagementBaseURL = "http://localhost:4100/accountmanagement"
	accountsBaseURL           = accountsManagementBaseURL + "/accounts"
)

type Item struct {
	ID   string
	Name string
}

func Index(req *restful.Request, resp *restful.Response) {
	root := hal.NewResource(struct{}{}, accountsBaseURL)
	root.AddLink("alternative", hal.NewLink(accountsBaseURL+"accounts?page=1"))
	root.AddLink("account", hal.NewLink(accountsBaseURL+"/{accountID}"))
	root.AddLink("account", hal.NewLink(accountsBaseURL+"/123"))
	root.AddLink("account", hal.NewLink(accountsBaseURL+"/456"))
	root.AddLink("next", hal.NewLink(accountsBaseURL+"?page=2"))
	root.AddLink("parent", hal.NewLink(accountsManagementBaseURL))

	resp.WriteHeaderAndJson(http.StatusOK, root, conf.MIME_HAL_JSON)
}

func Account(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("accountID")

	root := hal.NewResource(
		Item{ID: id, Name: "Account_" + id},
		accountsBaseURL+"/"+id)

	root.AddLink("collection", hal.NewLink(accountsBaseURL))
	root.AddLink("followers", hal.NewLink(fmt.Sprintf("%s/%s/followers", accountsBaseURL, id)))
	root.AddLink("following", hal.NewLink(fmt.Sprintf("%s/%s/following", accountsBaseURL, id)))
	root.AddLink("blurbs", hal.NewLink(fmt.Sprintf("%s/%s/blurbs", accountsBaseURL, id)))

	resp.WriteHeaderAndJson(http.StatusOK, root, conf.MIME_HAL_JSON)
}
