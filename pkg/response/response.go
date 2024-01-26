package response

type MetaResponse struct {
	Count int `json:"count"`
	Total int `json:"total"`
}

type HTTPResponse[T any] struct {
	Message string        `json:"message"`
	Result  *T            `json:"result"`
	Meta    *MetaResponse `json:"meta"`
}
