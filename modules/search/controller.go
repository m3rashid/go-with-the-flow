package search

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/m3rashid/go-server/db"
	search "github.com/m3rashid/go-server/modules/search/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleSearch() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		searchBody := struct {
			Text string `json:"text" validate:"required"`
		}{}
		if err := ctx.BodyParser(&searchBody); err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Could Not Parse Body")
		}

		collection := db.OpenCollection(db.Client, search.RESOURCE_MODEL_NAME)

		var results []search.Resource
		cursor, err := collection.Find(context.Background(), primitive.M{
			"$text": primitive.M{
				"$search": searchBody.Text,
			},
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Could Not Find Resources")
		}

		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		jsonResults, err := json.Marshal(results)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Could Not Marshal Results")
		}

		return ctx.Status(fiber.StatusOK).Send(jsonResults)
	}
}

func CreateResource(
	Name string,
	ResourceID primitive.ObjectID,
	ResourceType string,
	Description string,
) (err error) {
	collection := db.OpenCollection(db.Client, "search")
	newResource := search.Resource{
		Name:         Name,
		Description:  Description,
		ResourceID:   ResourceID,
		ResourceType: ResourceType,
	}

	_, err = collection.InsertOne(context.Background(), newResource)
	return
}
