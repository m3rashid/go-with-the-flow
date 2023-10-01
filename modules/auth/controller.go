package auth

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/m3rashid/server/db"
	auth "github.com/m3rashid/server/modules/auth/schema"
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
		collection := db.OpenCollection(db.Client, auth.USER_MODEL_NAME)

		newUser := auth.User{}
		err := ctx.BodyParser(&newUser)

		if err != nil {
			return ctx.Status(http.StatusBadRequest).SendString("Bad Request")
		}

		_, err = collection.InsertOne(ctx.Context(), newUser)
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).SendString("Could Not Register User, Please try again later")
		}

		return ctx.Status(http.StatusCreated).JSON(newUser)
	}
}