package request

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"project1/usecases/domain"

	"github.com/gorilla/mux"
)

func DecodeRequestPersonalProfileByID(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	idStr, ok := vars["personal_id"]
	if !ok || idStr == "" {
		return nil, fmt.Errorf("personal_id is required")
	}
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid personal id: %v", err)
	}
	return domain.PersonalProfile{Id: id}, nil
}


func DecodeRequestPersonalProfileAll(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}


func DecodeRequestPersonalProfilePost(_ context.Context, r *http.Request) (interface{}, error) {
	var req domain.PersonalProfile
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request body: %v", err)
	}
	return req, nil
}


func DecodeRequestPersonalProfilePatch(_ context.Context, r *http.Request) (interface{}, error) {
    vars := mux.Vars(r)
    idStr, ok := vars["personal_id"]
    if !ok || idStr == "" {
        return nil, fmt.Errorf("personal_id is required")
    }
    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        return nil, fmt.Errorf("invalid personal_id: %v", err)
    }

    var req domain.PersonalProfile
    if r.ContentLength > 0 {
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            return nil, fmt.Errorf("invalid request body: %v", err)
        }
    }
    req.Id = id
    return req, nil
}


func DecodeRequestPersonalProfileDelete(_ context.Context, r *http.Request) (interface{}, error) {
    vars := mux.Vars(r)
    idStr, ok := vars["personal_id"]
    if !ok || idStr == "" {
        return nil, fmt.Errorf("personal id is required")
    }

    id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        return nil, fmt.Errorf("invalid personal_id: %v", err)
    }

    return domain.PersonalProfile{Id: id}, nil
}

