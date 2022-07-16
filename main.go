package main

import (
	"os"
	"seior-shortener-url/database"
	"seior-shortener-url/handler"
	"seior-shortener-url/model/entity"
	"seior-shortener-url/repository"
	"seior-shortener-url/router"
	"seior-shortener-url/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := database.NewDB()

	// migrate the schema
	db.AutoMigrate(&entity.Url{})

	app := fiber.New(fiber.Config{
		ErrorHandler: handler.ErrorHandler,
	})

	// registering middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// registering repository
	urlRepository := repository.NewUrlRepository(db)

	// registering service
	urlService := service.NewUrlService(urlRepository)

	// registering handler
	urlHandler := handler.NewUrlHandler(urlService)

	// registering router
	router.BindUrlHandler(app, urlHandler)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Server Running !!!",
		})
	})

	err = app.Listen(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}

}
