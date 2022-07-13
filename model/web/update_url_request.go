package web

type UpdateUrlRequest struct {
	Id          string `json:"id"`
	Destination string `json:"destination"`
	Alias       string `json:"alias"`
}
