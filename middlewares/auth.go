package middlewares

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Email      string    `json:"email"`
	UserID     string    `json:"userId"`
	Exp        time.Time `json:"exp"`
	Authorized bool      `json:"authorized"`
	jwt.StandardClaims
}

func CheckAuth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		clientToken := ctx.Get("Authorization")
		if clientToken == "" {
			return ctx.Status(fiber.StatusInternalServerError).SendString("No Authorization header provided")
		}

		token, err := jwt.ParseWithClaims(clientToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

		claims, ok := token.Claims.(*Claims)
		if !ok {
			return ctx.Status(fiber.StatusInternalServerError).SendString("The token is invalid")
		}

		if !claims.Authorized || claims.Exp.Before(time.Now()) {
			return ctx.Status(fiber.StatusUnauthorized).SendString("The token is invalid")
		}

		ctx.Locals("email", claims.Email)
		ctx.Locals("userId", claims.UserID)
		ctx.Locals("authorized", claims.Authorized)

		return ctx.Next()
	}
}
