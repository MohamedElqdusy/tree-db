package service

import (
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Method string            //HTTP method
	Path   string            //url endpoint
	Handle httprouter.Handle //Controller function which dispatches the right HTML page and/or data for each route
}

type Routes []Route

var routes = Routes{
	Route{
		"GET",
		"/",
		Index,
	},
	Route{
		"GET",
		"/nodechilds/:node_id",
		FindNodeChilds,
	},
	Route{
		"PUT",
		"/update/:node_id/:parent_id",
		UpdateParent,
	},
	Route{
		"POST",
		"/node",
		StoreNode,
	},
}
