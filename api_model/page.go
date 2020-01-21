package api_model

type Page struct {
	Index int `json:"index"`
	Count int `json:"count"`
	Total int `json:"total"`
}
