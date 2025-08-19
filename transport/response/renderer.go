package response

import (
	"context"
	"encoding/json"
	"net/http"
	"project1/config"

	error_handler "project1/error-handler"
)

type errorResponse struct {
	Message string      `json:"message"`
	
}

func HandleError(ctx context.Context, err error, w http.ResponseWriter) {

	if !isPublicVisible(err) {
		errResponse := errorResponse{
			Message: `Something went wrong`,
		}
	
		genericError(ctx, errResponse, w)
		return
	}

	if domainError, ok := err.(*error_handler.DomainError); ok {

		body := errorResponse{
			Message: domainError.Message,
		}


		w.Header().Set("Content-Type", "application/json")
		 w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(body)
		return
	}

	if applicationError, ok := err.(*error_handler.ApplicationError); ok  {

		body := errorResponse{
			Message: applicationError.Message,
		}

		
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(applicationError.StatusCode)
		json.NewEncoder(w).Encode(body)
	}
}

func isPublicVisible(err error) bool {
	_, isDomain := err.(*error_handler.DomainError)
	_, isApplication := err.(*error_handler.ApplicationError)
	return isDomain || (config.AppConf.Debug && isApplication)
}


func isDomainError(err error) (*error_handler.DomainError, bool) {
	domainError, isDomain := err.(*error_handler.DomainError)
	return domainError, isDomain
}

func isApplicationError(err error) (*error_handler.ApplicationError, bool) {
	applicationError, isApplication := err.(*error_handler.ApplicationError)
	return applicationError, isApplication
}


func genericError(_ context.Context, err errorResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err)
}

