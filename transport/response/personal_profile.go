package response

import (
	"context"
	"encoding/json"
	"net/http"

	"project1/transport/endpoints"
)

func EncodeResponsePersonalProfile(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := response.(endpoints.PersonalProfileRequestResponse)

	result := res.Response

	//Response wrapper
	apiResponse := setResponse(result)

	return json.NewEncoder(w).Encode(apiResponse)
}

