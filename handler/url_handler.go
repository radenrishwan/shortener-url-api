package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"seior-shortener-url/model/web"
	"seior-shortener-url/service"
)

type UrlHandler interface {
	CreateUrl(ctx *fiber.Ctx) error
	UpdateUrl(ctx *fiber.Ctx) error
	DeleteUrl(ctx *fiber.Ctx) error
	GetUrlInfo(ctx *fiber.Ctx) error
	Redirect(ctx *fiber.Ctx) error
}

type urlHandler struct {
	service.UrlService
}

func NewUrlHandler(urlService service.UrlService) UrlHandler {
	return &urlHandler{UrlService: urlService}
}

func (handler *urlHandler) CreateUrl(ctx *fiber.Ctx) error {
	var request web.CreateUrlRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		panic(err)
	}

	result := handler.UrlService.Create(request)

	return ctx.Status(http.StatusOK).JSON(result)
}

func (handler *urlHandler) UpdateUrl(ctx *fiber.Ctx) error {
	var request web.UpdateUrlRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		panic(err)
	}

	result := handler.UrlService.Update(request)

	return ctx.Status(http.StatusOK).JSON(result)
}

func (handler *urlHandler) DeleteUrl(ctx *fiber.Ctx) error {
	var request web.DeleteUrlRequest

	request.Id = ctx.Query("id")

	result := handler.UrlService.Delete(request)

	return ctx.Status(http.StatusOK).JSON(result)
}

func (handler *urlHandler) GetUrlInfo(ctx *fiber.Ctx) error {
	var request web.FindByAliasUrlRequest

	request.Alias = ctx.Query("alias")

	result := handler.UrlService.FindByAlias(request)

	return ctx.Status(http.StatusOK).JSON(result)
}

func (handler *urlHandler) Redirect(ctx *fiber.Ctx) error {
	var request web.FindByAliasUrlRequest

	request.Alias = ctx.Params("alias")

	result := handler.UrlService.RedirectUrl(request)

	return ctx.Redirect(result.Data.Destination)
}
