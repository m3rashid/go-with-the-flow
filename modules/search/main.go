package search

import (
	"github.com/m3rashid/server/modules"
)

var SearchModule = modules.Module{
	Name: "search",
	AuthenticatedRoutes: modules.Controller{
		"/search": HandleSearch(),
	},
	AnonymousRoutes: modules.Controller{},
}
