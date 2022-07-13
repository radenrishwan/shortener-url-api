package router

import (
	"github.com/gofiber/fiber/v2"
	"seior-shortener-url/handler"
)

func BindUrlHandler(app *fiber.App, handler handler.UrlHandler) {
	app.Post("/api/url", handler.CreateUrl)
	app.Delete("/api/url/", handler.DeleteUrl)
	app.Get("/api/url/", handler.GetUrlInfo)
	app.Put("/api/url/", handler.UpdateUrl)

	app.Get("/:alias", handler.Redirect)
}
