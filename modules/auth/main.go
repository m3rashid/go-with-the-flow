package auth

import (
	"github.com/m3rashid/go-with-the-flow/modules"
	auth "github.com/m3rashid/go-with-the-flow/modules/auth/schema"
	search "github.com/m3rashid/go-with-the-flow/modules/search/schema"
)

var AuthModule = modules.Module{
	Name: "auth",
	Permissions: []modules.ModulePermission{
		{
			Name:         "user",
			ResourceType: auth.USER_MODEL_NAME,
			ResourceIndex: search.ResourceIndex{
				NameKey:        "name",
				DescriptionKey: "email",
				DisplayUrl:     "/user/:rId",
			},
			ActionPermissions:      modules.Permission{},
			IndependentPermissions: modules.Permission{},
		},
		{
			Name:          "profile",
			ResourceType:  auth.PROFILE_MODEL_NAME,
			ResourceIndex: search.ResourceIndex{},
			ActionPermissions: modules.Permission{
				"CAN_VIEW_PROFILE":   {},
				"CAN_EDIT_PROFILE":   {},
				"CAN_DELETE_PROFILE": {},
			},
			IndependentPermissions: modules.Permission{
				"CAN_CREATE_PROFILE": {},
			},
		},
	},
	AuthenticatedRoutes: modules.Controller{},
	AnonymousRoutes: modules.Controller{
		"/login":    Login(),
		"/register": Register(),
	},
}
