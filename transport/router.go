package http

import (
	"fmt"
	"log"
	"net/http"
	"project1/config"

	"project1/services"
	"project1/transport/endpoints"
	httpRequest "project1/transport/request"
	"project1/transport/response"
	httpResponse "project1/transport/response"

	kitLogger "github.com/go-kit/kit/log"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"os"
)

func Listen() {
	mux := Handlers()
	port := config.AppConf.Port
	err := http.ListenAndServe(fmt.Sprintf(`:%d`, port), mux)
	if err != nil {
		log.Fatal(`Cannot start web server : `, err)
	}
}

func Handlers() *mux.Router {
	r := mux.NewRouter()

	//kitLogger initialization
	logger := kitLogger.NewLogfmtLogger(os.Stderr)
	opts := []httpTransport.ServerOption{
		httpTransport.ServerErrorLogger(logger),
		httpTransport.ServerErrorEncoder(response.HandleError),
	}

	var personalProfileService services.PersonalProfileService
	personalProfileService = services.PersonalProfile{}

	r.Handle(`/personal/profile/{personal_id}`, httpTransport.NewServer(
		endpoints.PersonalProfileEndpoint(personalProfileService),
		httpRequest.DecodeRequestPersonalProfileByID,
		httpResponse.EncodeResponsePersonalProfile,
		opts...,
	)).Methods(http.MethodGet)

	r.Handle(`/personal/profile`, httpTransport.NewServer(
		endpoints.GetAllPersonalProfilesEndpoint(personalProfileService),
		httpRequest.DecodeRequestPersonalProfileAll,
		httpResponse.EncodeResponsePersonalProfile,
		opts...,
	)).Methods(http.MethodGet)

	r.Handle(`/personal/profile`, httpTransport.NewServer(
		endpoints.CreatePersonalProfileEndpoint(personalProfileService),
		httpRequest.DecodeRequestPersonalProfilePost, 
		httpResponse.EncodeResponsePersonalProfile,
		opts...,
	)).Methods(http.MethodPost)

	r.Handle(`/personal/profile/{personal_id}`, httpTransport.NewServer(
	endpoints.UpdatePersonalProfileEndpoint(personalProfileService),
	httpRequest.DecodeRequestPersonalProfilePatch,
	httpResponse.EncodeResponsePersonalProfile,
	opts...,
)).Methods(http.MethodPatch)

	//Metrics handler
	r.Handle(`/metrics`, promhttp.Handler())

	return r
}
