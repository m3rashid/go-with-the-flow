package search

import (
	"github.com/m3rashid/go-with-the-flow/modules"
)

var SearchModule = modules.Module{
	Name: "search",
	AuthenticatedRoutes: modules.Controller{
		"/search": HandleSearch(),
	},
	AnonymousRoutes: modules.Controller{},
}
