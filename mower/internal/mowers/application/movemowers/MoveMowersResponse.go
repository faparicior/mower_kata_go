package movemowers

type MoveMowersResponse struct {
	response string
}

func BuildMoveMowersResponse(response string) MoveMowersResponse {
	return MoveMowersResponse{response: response}
}

func (response MoveMowersResponse) Response() string {
	return response.response
}
