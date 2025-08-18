package usecases

import (
	"context"
	error_handler "project1/error-handler"
	"time"

	log "project1/logger"
	"project1/usecases/domain"
)

type PersonalProfileInterface struct {
	PersonalProfileRepository domain.PersonalProfileRepository
}

func (interactor PersonalProfileInterface) GetPersonalProfile(ctx context.Context, request domain.PersonalProfile) (response []domain.PersonalProfile, err error) {
	result, err := interactor.PersonalProfileRepository.GetPersonalProfile(ctx, request)
	if err != nil {
		err = error_handler.ApplicationError{Message: `Something went wrong'`, Details: err}
		log.Error(log.WithPrefix("personal-services:personal-api", "GetPersonalProfile"), err)
		return result, err
	}

	if len(result) == 0 {
		err = error_handler.ApplicationError{Message: `Something went wrong'`, Details: err}
		log.Debug(log.WithPrefix("personal-services:personal-api", "GetPersonalProfile"), request.Id, err)
		return result, err
	}

	return result, err
}

func (interactor PersonalProfileInterface) CreatePersonalProfile(ctx context.Context, request domain.PersonalProfile) (response domain.PersonalProfile, err error) {
	now := time.Now()
	request.CreateTime = now
	request.UpdateTime = now

	result, err := interactor.PersonalProfileRepository.CreatePersonalProfile(ctx, request)
	if err != nil {
		err = error_handler.ApplicationError{Message: "Failed to create personal profile", Details: err}
		log.Error(log.WithPrefix("personal-services:personal-api", "CreatePersonalProfile"), err)
		return result, err
	}

	log.Info(log.WithPrefix("personal-services:personal-api", "CreatePersonalProfile"), "Created profile ID:", result.Id)
	return result, nil
}


func (interactor PersonalProfileInterface) GetAllPersonalProfiles(ctx context.Context) (response []domain.PersonalProfile, err error) {
	result, err := interactor.PersonalProfileRepository.GetAllPersonalProfiles(ctx)
	if err != nil {
		err = error_handler.ApplicationError{Message: "Failed to fetch all personal profiles", Details: err}
		log.Error(log.WithPrefix("personal-services:personal-api", "GetAllPersonalProfiles"), err)
		return nil, err
	}

	if len(result) == 0 {
		log.Debug(log.WithPrefix("personal-services:personal-api", "GetAllPersonalProfiles"), "No profiles found")
	}

	return result, nil
}

func (interactor PersonalProfileInterface) UpdatePersonalProfile(ctx context.Context, request domain.PersonalProfile) (response domain.PersonalProfile, err error) {
	request.UpdateTime = time.Now()
	result, err := interactor.PersonalProfileRepository.UpdatePersonalProfile(ctx, request)
	if err != nil {
		err = error_handler.ApplicationError{Message: "Failed to update personal profile", Details: err}
		log.Error(log.WithPrefix("personal-services:personal-api", "UpdatePersonalProfile"), err)
		return result, err
	}
	log.Info(log.WithPrefix("personal-services:personal-api", "UpdatePersonalProfile"), "Updated profile ID:", result.Id)
	return result, nil
}
