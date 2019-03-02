package controllers

import (
	restful "github.com/emicklei/go-restful"
)

const (
	entryPointBaseURL = "http://localhost:4100/"
	accountsBaseURL   = "http://localhost:4101/accountmanagement/"
)

func Index(req *restful.Request, resp *restful.Response) {
	resp.Header().Set("Content-Type", "application/hal+json")

	/*
		root := hal.NewResource(struct{}{}, entryPointBaseURL+"/accountmanagement")
		root.AddLink("accounts", hal.NewLink(accountsBaseURL+"/accountmanagement/accounts"))

		data, _ := json.Marshal(root)
		resp.Write(data)
	*/
}

func Account(req *restful.Request, resp *restful.Response) {
	resp.Header().Set("Content-Type", "application/hal+json")

	/*
		root := hal.NewResource(struct{}{}, entryPointBaseURL+"/accountmanagement")
		root.AddLink("accounts", hal.NewLink(accountsBaseURL+"/accountmanagement/accounts"))

		data, _ := json.Marshal(root)
		resp.Write(data)
	*/
}
