package controllers

import (
	"fmt"
	"net/http"

	"github.com/sammyne/ppp-ddd/chapter13/app-restful/events"

	restful "github.com/emicklei/go-restful"
	"github.com/nvellon/hal"
	"github.com/sammyne/ppp-ddd/chapter13/app-restful/conf"
)

type FollowerBrief struct {
	AccountID string
}

type FollowersList struct {
	Followers []*FollowerBrief
}

func Followers(req *restful.Request, resp *restful.Response) {
	fmt.Println("hello world")
	id := req.PathParameter("accountID")
	href := accountsBaseURL + "/" + id + "/followers"

	root := hal.NewResource(FollowersList{Followers: dummyFollowerBriefs(id)}, href)
	root.AddNewLink("next", href+"?pages=2")

	resp.WriteHeaderAndJson(http.StatusOK, root, conf.MIME_HAL_JSON)
}

func Follow(req *restful.Request, resp *restful.Response) {
	// accountId will be taken from querystring - it is a simple type
	// follower will be taken from request body - it is a complex type
	accountID := req.PathParameter("accountID")
	followerID := req.QueryParameter("followerID")

	//fmt.Println(accountID, followerID)

	event := events.NewBeganFollowing(accountID, followerID)
	events.Persist(event)

	to := fmt.Sprintf("%s/%s/followers", accountsBaseURL, accountID)
	http.Redirect(resp.ResponseWriter, req.Request, to, http.StatusSeeOther)
}

func dummyFollowerBriefs(id string) []*FollowerBrief {
	return []*FollowerBrief{{"follower1"}, {"follower2"}, {"follower3"}}
}
