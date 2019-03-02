package regular

import (
	restful "github.com/emicklei/go-restful"
	"github.com/sammyne/ppp-ddd/chapter13/app-restful/events"
)

const baseURL = "http://localhost:4102/"

func Feeds(req *restful.Request, resp *restful.Response) {
	events.Retrieve(nil)
}
