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
	authSchema "github.com/m3rashid/go-with-the-flow/modules/auth/schema"
	"github.com/m3rashid/go-with-the-flow/modules/flow"
	searchSchema "github.com/m3rashid/go-with-the-flow/modules/search/schema"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New(fiber.Config{
		CaseSensitive:  true,
		AppName:        os.Getenv("APP_NAME"),
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

	go func() {
		flow.StartWatchMongo([]string{
			authSchema.USER_MODEL_NAME,
			authSchema.PROFILE_MODEL_NAME,
			searchSchema.RESOURCE_MODEL_NAME,
		})
	}()

	log.Println("Server is running")
	app.Listen(":" + os.Getenv("SERVER_PORT"))
}
