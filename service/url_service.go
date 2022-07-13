package service

import (
	"github.com/google/uuid"
	"net/http"
	"seior-shortener-url/exception"
	"seior-shortener-url/helper"
	"seior-shortener-url/model/entity"
	"seior-shortener-url/model/web"
	"seior-shortener-url/repository"
	"seior-shortener-url/validation"
	"strings"
)

type UrlService interface {
	Create(request web.CreateUrlRequest) web.DefaultResponse[web.UrlResponse]
	Update(request web.UpdateUrlRequest) web.DefaultResponse[web.UrlResponse]
	Delete(request web.DeleteUrlRequest) web.DefaultResponse[web.UrlResponse]
	FindByAlias(request web.FindByAliasUrlRequest) web.DefaultResponse[web.UrlResponse]
	RedirectUrl(request web.FindByAliasUrlRequest) web.DefaultResponse[web.UrlResponse]
}

type urlService struct {
	repository.UrlRepository
}

func NewUrlService(urlRepository repository.UrlRepository) UrlService {
	return &urlService{UrlRepository: urlRepository}
}

func (service *urlService) Create(request web.CreateUrlRequest) web.DefaultResponse[web.UrlResponse] {
	validation.NewCreateUrlValidation(request)
	time := helper.GenerateMilisTimeNow()

	url, err := service.UrlRepository.Create(entity.Url{
		Id:          uuid.NewString(),
		Destination: request.Destination,
		Alias:       request.Alias,
		Clicked:     0,
		CreatedAt:   time,
		UpdatedAt:   time,
	})

	if err != nil {
		panic(exception.NewIsExistException(err.Error()))
	}

	return web.DefaultResponse[web.UrlResponse]{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    web.NewUrlResponse(url),
	}
}

func (service *urlService) Update(request web.UpdateUrlRequest) web.DefaultResponse[web.UrlResponse] {
	validation.NewUpdateUrlValidation(request)

	// find if exist
	url, err := service.UrlRepository.FindById(entity.Url{Id: request.Id})
	if err != nil {
		panic(exception.NewNotFoundException("Url not found"))
	}

	// edit property
	time := helper.GenerateMilisTimeNow()

	url.Alias = request.Alias
	url.Destination = request.Destination
	url.UpdatedAt = time

	// update
	service.UrlRepository.Update(url)

	return web.DefaultResponse[web.UrlResponse]{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    web.NewUrlResponse(url),
	}
}

func (service *urlService) Delete(request web.DeleteUrlRequest) web.DefaultResponse[web.UrlResponse] {
	validation.NewDeleteUrlValidation(request)

	// find if exist
	url, err := service.UrlRepository.FindById(entity.Url{Id: request.Id})
	if err != nil {
		panic(exception.NewNotFoundException("Url not found"))
	}

	service.UrlRepository.Delete(url)

	return web.DefaultResponse[web.UrlResponse]{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    web.NewUrlResponse(url),
	}
}

func (service *urlService) FindByAlias(request web.FindByAliasUrlRequest) web.DefaultResponse[web.UrlResponse] {
	validation.NewFindByAliasUrlValidation(request)

	url, err := service.UrlRepository.FindByAlias(entity.Url{Alias: request.Alias})
	if err != nil {
		panic(exception.NewNotFoundException("Url not found"))
	}

	service.UrlRepository.Update(url)

	return web.DefaultResponse[web.UrlResponse]{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    web.NewUrlResponse(url),
	}
}

func (service *urlService) RedirectUrl(request web.FindByAliasUrlRequest) web.DefaultResponse[web.UrlResponse] {
	validation.NewFindByAliasUrlValidation(request)

	url, err := service.UrlRepository.FindByAlias(entity.Url{Alias: request.Alias})
	if err != nil {
		panic(exception.NewNotFoundException("Url not found"))
	}

	// update click count
	url.Clicked += 1
	service.UrlRepository.Update(url)

	if !strings.Contains(url.Destination, "http") {
		url.Destination = "https://" + url.Destination
	}

	return web.DefaultResponse[web.UrlResponse]{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    web.NewUrlResponse(url),
	}
}
