package web

type CreateUrlRequest struct {
	Destination string `json:"destination"`
	Alias       string `json:"alias"`
}
