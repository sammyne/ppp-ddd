package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsersFollowers(resp http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(dummyFollowers())
	if nil != err {
		resp.WriteHeader(http.StatusInternalServerError)
	}

	resp.Write(data)
}

func dummyFollowers() []*Follower {
	followers := make([]*Follower, 10)
	for i := range followers {
		followers[i] = &Follower{
			ID:         fmt.Sprintf("follower_%d", i),
			Name:       fmt.Sprintf("happy follower %d", i),
			SocialTags: []string{"programming", "DDD", "Psychology"},
		}
	}

	return followers
}
