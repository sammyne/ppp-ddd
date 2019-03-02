package controllers

import (
	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/nvellon/hal"
)

const (
	entryPointBaseURL = "http://localhost:4100"
	accountsBaseURL   = "http://localhost:4101"
)

func GetEntryPoint(req *restful.Request, resp *restful.Response) {
	resp.Header().Set("Content-Type", "application/hal+json")
	//resp.Write([]byte("hello world"))

	root := hal.NewResource(struct{}{}, entryPointBaseURL+"/accountmanagement")
	root.AddLink("accounts", hal.NewLink(accountsBaseURL+"/accountmanagement/accounts"))

	//resp.WriteAsJson(root)
	data, _ := json.Marshal(root)
	resp.Write(data)
}
