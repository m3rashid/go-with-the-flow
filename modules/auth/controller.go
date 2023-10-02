package auth

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/m3rashid/server/db"
	auth "github.com/m3rashid/server/modules/auth/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		collection := db.OpenCollection(db.Client, auth.USER_MODEL_NAME)

		loginBody := struct {
			Email    string `json:"email" validate:"required,email"`
			Password string `json:"password" validate:"required"`
		}{}
		err := ctx.BodyParser(&loginBody)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).SendString("Could Not Parse Body")
		}

		validate := validator.New()
		err = validate.Struct(loginBody)
		if err != nil {
			return ctx.Status(http.StatusBadRequest).SendString("Validation Error")
		}

		var user auth.User
		err = collection.FindOne(ctx.Context(), auth.User{Email: loginBody.Email}).Decode(&user)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		return ctx.Status(http.StatusOK).JSON(user)
	}
}

func Register() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("Registering User")

		collection := db.OpenCollection(db.Client, auth.USER_MODEL_NAME)

		newUser := auth.User{}
		err := ctx.BodyParser(&newUser)
		if err != nil {
			log.Println(err)
			return ctx.Status(http.StatusBadRequest).SendString("Bad Request")
		}

		newUser.Deactivated = false
		newUser.Deleted = false
		newUser.Roles = []primitive.ObjectID{}

		validator := validator.New()
		err = validator.Struct(newUser)

		if err != nil {
			log.Println(err)
			return ctx.Status(http.StatusBadRequest).SendString("Bad Request")
		}

		res, err := collection.InsertOne(ctx.Context(), newUser)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString("Could Not Register User, Please try again later")
		}

		return ctx.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "User Registered Successfully",
			"userId":  res.InsertedID,
		})
	}
}
