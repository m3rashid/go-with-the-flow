package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(userId string, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["userId"] = userId
	claims["email"] = email

	return token.SignedString(os.Getenv("JWT_SECRET"))
}
