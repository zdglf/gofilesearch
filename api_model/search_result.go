package api_model

type SearchResult struct {
	Id       string   `json:"id"`
	Url      string   `json:"url"`
	Name     string   `json:"name"`
	Desc     []string `json:"desc"`
	CreateAt string   `json:"createAt"`
}

type SearchResultResponse struct {
	*Page
	*StatusCode
	Data []SearchResult `json:"data"`
}
