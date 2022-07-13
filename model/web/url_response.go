package web

import "seior-shortener-url/model/entity"

type UrlResponse struct {
	Id          string `json:"id"`
	Destination string `json:"destination"`
	Alias       string `json:"alias"`
	Clicked     uint64 `json:"clicked"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewUrlResponse(url entity.Url) UrlResponse {
	return UrlResponse{
		Id:          url.Id,
		Destination: url.Destination,
		Alias:       url.Alias,
		Clicked:     url.Clicked,
		CreatedAt:   url.CreatedAt,
		UpdatedAt:   url.UpdatedAt,
	}
}
