package response

type Response struct {
	Data interface{} `json:"data"`
}

func setResponse(data interface{}) (res Response) {
	response := Response{
		Data: data,
	}
	return response
}
