package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Users(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	// the accountID isn't effective yet
	url := "http://localhost:3200/api/followerdirectory/getusersfollowers?accountID=" + req.PostFormValue("accountID")

	response, err := http.Get(url)
	if nil != err {
		fmt.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var followers []*Follower
	if err := json.NewDecoder(response.Body).Decode(&followers); nil != err {
		fmt.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	usersJSON, err := json.Marshal(findUsersBySocialTags(followers))
	if nil != err {
		fmt.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp.Write(usersJSON)
}

func findUsersBySocialTags(followers []*Follower) []string {
	tags := make(map[string]struct{})

	// flatten the tags and compose them as a set
	/*
	* Real system would look at the users tags and find
	* popular accounts with similar tags by making more
	* RPC calls.
	 */
	for _, follower := range followers {
		for _, tag := range follower.SocialTags {
			tags[tag] = struct{}{}
		}
	}

	out := make([]string, 0, len(tags))
	for tag := range tags {
		out = append(out, tag+"_user_1")
	}

	return out
}
