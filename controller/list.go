package controller

import (
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	CollectionName string
	DbStruct       interface{}
	SearchQuery    map[string]interface{}
}

func List(c Controller) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		// collection := db.OpenCollection(db.Client, c.CollectionName)

		return nil
	}
}
