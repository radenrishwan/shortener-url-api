package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
	"seior-shortener-url/exception"
	"seior-shortener-url/model/web"
)

func NewCreateUrlValidation(request web.CreateUrlRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Destination,
			validation.Required,
			validation.Match(regexp.MustCompile("[(http(s)?):\\/\\/(www\\.)?a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&\\\\=]*)")).
				Error("Destination must be valid URL")),
		validation.Field(&request.Alias, validation.Required, validation.Length(3, 255)),
	)

	if err != nil {
		panic(exception.NewValidationException(err.Error()))
	}
}

func NewUpdateUrlValidation(request web.UpdateUrlRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Destination,
			validation.Required,
			validation.Match(regexp.MustCompile("[(http(s)?):\\/\\/(www\\.)?a-zA-Z0-9@:%._\\+~#=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%_\\+.~#?&\\\\=]*)")).
				Error("Destination must be valid URL")),
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Alias, validation.Required, validation.Length(3, 255)),
	)

	if err != nil {
		panic(exception.NewValidationException(err.Error()))
	}
}

func NewDeleteUrlValidation(request web.DeleteUrlRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
	)

	if err != nil {
		panic(exception.NewValidationException(err.Error()))
	}
}

func NewFindByAliasUrlValidation(request web.FindByAliasUrlRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Alias, validation.Required),
	)

	if err != nil {
		panic(exception.NewValidationException(err.Error()))
	}
}
