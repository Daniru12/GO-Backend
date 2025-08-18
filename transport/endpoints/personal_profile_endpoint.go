package endpoints

import (
	"context"

	"project1/services"
	"project1/usecases/domain"

	"github.com/go-kit/kit/endpoint"
)

type PersonalProfileRequestResponse struct {
	Request  domain.PersonalProfile
	Response interface{}
}

func PersonalProfileEndpoint(svc services.PersonalProfileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		requestStatus, _ := request.(domain.PersonalProfile)

		responseStatus, err := svc.GetPersonalProfile(ctx, requestStatus)

		return PersonalProfileRequestResponse{
			requestStatus,
			responseStatus,
		}, err
	}
}

func GetAllPersonalProfilesEndpoint(svc services.PersonalProfileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// No input needed, fetch all
		responseStatus, err := svc.GetAllPersonalProfiles(ctx)
		if err != nil {
			return nil, err
		}

		return PersonalProfileRequestResponse{
			Request:  domain.PersonalProfile{}, // empty request
			Response: responseStatus,
		}, nil
	}
}



func CreatePersonalProfileEndpoint(svc services.PersonalProfileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		requestStatus, _ := request.(domain.PersonalProfile)

		responseStatus, er := svc.CreatePersonalProfile(ctx, requestStatus)

		return PersonalProfileRequestResponse{
			Request:  requestStatus,
			Response: responseStatus,
		}, er
	}
}

func UpdatePersonalProfileEndpoint(svc services.PersonalProfileService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, _ := request.(domain.PersonalProfile)
		res, err := svc.UpdatePersonalProfile(ctx, req)
		return PersonalProfileRequestResponse{
			Request:  req,
			Response: res,
		}, err
	}
}
