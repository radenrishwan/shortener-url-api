package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"seior-shortener-url/exception"
	"seior-shortener-url/model/web"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(exception.ValidationException)
	if ok {
		return ctx.Status(http.StatusBadRequest).JSON(web.DefaultResponse[string]{
			Code:    http.StatusBadRequest,
			Message: "Validation error",
			Data:    err.Error(),
		})
	}

	_, ok = err.(exception.NotFoundException)
	if ok {
		return ctx.Status(http.StatusNotFound).JSON(web.DefaultResponse[string]{
			Code:    http.StatusNotFound,
			Message: "Nof Found",
			Data:    err.Error(),
		})
	}

	_, ok = err.(exception.IsExistException)
	if ok {
		return ctx.Status(http.StatusConflict).JSON(web.DefaultResponse[string]{
			Code:    http.StatusBadRequest,
			Message: "Validation error",
			Data:    err.Error(),
		})
	}

	return ctx.Status(http.StatusInternalServerError).JSON(web.DefaultResponse[string]{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Data:    err.Error(),
	})
}
