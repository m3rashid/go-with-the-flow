package modules

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/m3rashid/go-server/middlewares"
	auth "github.com/m3rashid/go-server/modules/auth/schema"
	"github.com/m3rashid/go-server/modules/flow"
	search "github.com/m3rashid/go-server/modules/search/schema"
)

type Permission = map[string]struct {
	Level int `json:"level"`
	Scope int `json:"scope"`
}

type ModulePermission struct {
	Name                   string               `json:"name"`
	ResourceType           string               `json:"resourceType"`
	ResourceIndex          search.ResourceIndex `json:"resourceIndex"`
	ActionPermissions      Permission           `json:"actionPermissions"`
	IndependentPermissions Permission           `json:"independentPermissions"`
}

type Controller = map[string]fiber.Handler

type Module struct {
	Name                string             `json:"name"`
	Permissions         []ModulePermission `json:"permissions"`
	AnonymousRoutes     Controller
	AuthenticatedRoutes Controller
}

func RegisterRoutes(app *fiber.App, modules []Module) {
	flow.StartWatchMongo([]string{
		auth.USER_MODEL_NAME,
		auth.PROFILE_MODEL_NAME,
		search.RESOURCE_MODEL_NAME,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/metrics", monitor.New(monitor.Config{
		Title: os.Getenv("APP_NAME") + " - Metrics",
	}))

	for _, module := range modules {
		for route, handler := range module.AuthenticatedRoutes {
			app.Post("/api/"+module.Name+route, middlewares.CheckAuth(), handler)
		}

		for route, handler := range module.AnonymousRoutes {
			app.Post("/api/anonymous/"+module.Name+route, handler)
		}
	}
}
