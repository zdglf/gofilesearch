package api_model

type DataResponse struct {
	*StatusCode
	Data interface{} `json:"data"`
}

type WithPageDataResponse struct {
	*StatusCode
	*Page
	Data interface{} `json:"data"`
}
