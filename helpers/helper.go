package helpers

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func APIResponse(code int, status string, message string, data interface{}) Response {
	meta := Meta{}
	meta.Code = code
	meta.Status = status
	meta.Message = message

	response := Response{}
	response.Meta = meta
	response.Data = data

	return response
}
