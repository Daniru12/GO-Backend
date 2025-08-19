package usecases

import (
	"context"
	error_handler "project1/error-handler"
	"time"
	"project1/util"
	log "project1/logger"
	"project1/usecases/domain"
)

type PersonalProfileInterface struct {
	PersonalProfileRepository domain.PersonalProfileRepository
}

func (interactor PersonalProfileInterface) GetPersonalProfile(ctx context.Context, request domain.PersonalProfile) (response []domain.PersonalProfile, err error) {
	result, err := interactor.PersonalProfileRepository.GetPersonalProfile(ctx, request)
	if err != nil {
		err = error_handler.NewApplicationError("User Id can't find", 404,1001, nil)
		log.Error(log.WithPrefix("personal-services:personal-api", "GetPersonalProfile"), err)
		return result, err
	}

	if len(result) == 0 {
		err = error_handler.NewApplicationError("User Id can't find", 404,1002, nil)
		return result, err
	}

	return result, err
}

func (interactor PersonalProfileInterface) CreatePersonalProfile(ctx context.Context, request domain.PersonalProfile) (response domain.PersonalProfile, err error) {
	now := utils.CustomTime{Time: time.Now()}

request.CreateTime = now
request.UpdateTime = now


	result, err := interactor.PersonalProfileRepository.CreatePersonalProfile(ctx, request)
	if err != nil {
		err = error_handler.NewApplicationError("Failed to create personal profile", 400,1003, nil)
		log.Error(log.WithPrefix("personal-services:personal-api", "CreatePersonalProfile"), err)
		return result, err
	}

	log.Info(log.WithPrefix("personal-services:personal-api", "CreatePersonalProfile"), "Created profile ID:", result.Id)
	return result, nil
}


func (interactor PersonalProfileInterface) GetAllPersonalProfiles(ctx context.Context) (response []domain.PersonalProfile, err error) {
	result, err := interactor.PersonalProfileRepository.GetAllPersonalProfiles(ctx)
	if err != nil {
		err = error_handler.NewApplicationError("Failed to fetch all personal profiles", 500,1004, nil)
		log.Error(log.WithPrefix("personal-services:personal-api", "GetAllPersonalProfiles"), err)
		return nil, err
	}

	if len(result) == 0 {
		log.Debug(log.WithPrefix("personal-services:personal-api", "GetAllPersonalProfiles"), "No profiles found")
	}

	return result, nil
}
func (interactor PersonalProfileInterface) UpdatePersonalProfile(ctx context.Context, request domain.PersonalProfile) (response domain.PersonalProfile, err error) {

    existingProfiles, err := interactor.PersonalProfileRepository.GetPersonalProfile(ctx, domain.PersonalProfile{Id: request.Id})
    if err != nil {
        err = error_handler.NewApplicationError("Failed to fetch existing profile", 500,1005, nil)

        log.Error(log.WithPrefix("personal-services:personal-api", "UpdatePersonalProfile"), err)
        return response, err
    }

    if len(existingProfiles) == 0 {
        err = error_handler.NewApplicationError("Profile not found", 404,1006, nil)

        log.Error(log.WithPrefix("personal-services:personal-api", "UpdatePersonalProfile"), err)
        return response, err
    }

    existing := existingProfiles[0]

    if request.Name != nil {
        existing.Name = request.Name
    }
    if request.Description != nil {
        existing.Description = request.Description
    }
    if request.Status != nil {
        existing.Status = request.Status
    }

    existing.UpdateTime = utils.CustomTime{Time: time.Now()}

    updated, err := interactor.PersonalProfileRepository.UpdatePersonalProfile(ctx, existing)
    if err != nil {
        err = error_handler.NewApplicationError("Check again user id", 404,1007, nil)

        log.Error(log.WithPrefix("personal-services:personal-api", "UpdatePersonalProfile"), err)
        return updated, err
    }

    log.Info(log.WithPrefix("personal-services:personal-api", "UpdatePersonalProfile"), "Updated profile ID:", updated.Id)
    return updated, nil
}


func (interactor PersonalProfileInterface) DeletePersonalProfile(ctx context.Context, request domain.PersonalProfile) (interface{}, error) {
   
    profiles, err := interactor.PersonalProfileRepository.GetPersonalProfile(ctx, domain.PersonalProfile{Id: request.Id})
    if err != nil {
        log.Error(log.WithPrefix("personal services:personal api", "DeletePersonalProfile"), err)
        return nil, error_handler.NewApplicationError("Failed to fetch profile", 500, 1008, nil)
    }

    if len(profiles) == 0 {
        errMsg := "Profile not found, Failed Delete"
        log.Error(log.WithPrefix("personal services:personal api", "DeletePersonalProfile"), errMsg)
        return nil, error_handler.NewApplicationError(errMsg, 404, 1009, nil)
    }

    profile := profiles[0]
    deletedStatus := "D"
    profile.Status = &deletedStatus
    profile.UpdateTime = utils.CustomTime{Time: time.Now()}

    _, err = interactor.PersonalProfileRepository.UpdatePersonalProfile(ctx, profile)
    if err != nil {
        log.Error(log.WithPrefix("personal services:personal api", "DeletePersonalProfile"), err)
        return nil, error_handler.NewApplicationError("Failed to delete profile, User Not Found, Check Id again", 404, 1010, nil)
    }

    log.Info(log.WithPrefix("personal services:personal api", "DeletePersonalProfile"), "deleted profile ID:", profile.Id)


    return map[string]string{"message": "Profile deleted successfully"}, nil
}
