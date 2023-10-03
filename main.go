package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/m3rashid/go-with-the-flow/modules"
	"github.com/m3rashid/go-with-the-flow/modules/auth"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		AppName:        os.Getenv("APP_NAME"),
		StrictRouting:  true,
		CaseSensitive:  true,
		RequestMethods: []string{"GET", "POST", "HEAD", "OPTIONS"},
	})

	app.Static("/public", "./public", fiber.Static{
		MaxAge:        3600,
		CacheDuration: 10 * time.Second,
	})

	app.Use(limiter.New(limiter.Config{
		Max:               60,
		Expiration:        1 * time.Minute,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Use(logger.New())

	modules.RegisterRoutes(app, []modules.Module{
		auth.AuthModule,
	})

	app.Listen(":" + os.Getenv("SERVER_PORT"))
}
