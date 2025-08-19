package services

import (
	"context"
	"project1/repositories"
	"project1/usecases"
	"project1/error-handler"
	"project1/usecases/domain"
)

type PersonalProfileService interface {
	GetPersonalProfile(ctx context.Context, req domain.PersonalProfile) (res []domain.PersonalProfile, err error)
	CreatePersonalProfile(ctx context.Context, req domain.PersonalProfile) (res domain.PersonalProfile, err error)
	GetAllPersonalProfiles(ctx context.Context) (res []domain.PersonalProfile, err error)
	UpdatePersonalProfile(ctx context.Context, req domain.PersonalProfile) (domain.PersonalProfile, error)
	DeletePersonalProfile(ctx context.Context, req domain.PersonalProfile) (interface{}, error)
}

type PersonalProfile struct{}

func (PersonalProfile) GetPersonalProfile(ctx context.Context, req domain.PersonalProfile) (res []domain.PersonalProfile, err error) {

	//Usecase Interactor & Repository
	personalProfileInterface := usecases.PersonalProfileInterface{}
	personalProfileInterface.PersonalProfileRepository = repositories.PersonalProfileRepository

	res, err = personalProfileInterface.GetPersonalProfile(ctx, req)

	return
}

func (PersonalProfile) GetAllPersonalProfiles(ctx context.Context) (res []domain.PersonalProfile, err error) {
	interactor := usecases.PersonalProfileInterface{
		PersonalProfileRepository: repositories.PersonalProfileRepository,
	}
	return interactor.GetAllPersonalProfiles(ctx)
}


func (PersonalProfile) CreatePersonalProfile(ctx context.Context, req domain.PersonalProfile) (res domain.PersonalProfile, err error) {
	personalProfileInterface := usecases.PersonalProfileInterface{
		PersonalProfileRepository: repositories.PersonalProfileRepository,
	}

	res, err = personalProfileInterface.CreatePersonalProfile(ctx, req)
	return
}

func (p PersonalProfile) UpdatePersonalProfile(ctx context.Context, req domain.PersonalProfile) (res domain.PersonalProfile, err error) {
    interactor := usecases.PersonalProfileInterface{
        PersonalProfileRepository: repositories.PersonalProfileRepository,
    }

    res, err = interactor.UpdatePersonalProfile(ctx, req)
    if err != nil {
   
        return res, error_handler.NewApplicationError("Check again user id", 404,0, nil)
    }

    return res, nil
}


func (p PersonalProfile) DeletePersonalProfile(ctx context.Context, req domain.PersonalProfile) (interface{}, error) {
	interactor := usecases.PersonalProfileInterface{
		PersonalProfileRepository: repositories.PersonalProfileRepository,
	}

	_, err := interactor.DeletePersonalProfile(ctx, req)
	if err != nil {
		return nil, error_handler.NewApplicationError("Failed to delete profile, User Not Found, Check Id again", 404, 1010, nil)
	}

	return map[string]string{
		"message": "Profile deleted successfully",
	}, nil
}
