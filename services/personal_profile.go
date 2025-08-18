package services

import (
	"context"
	"project1/repositories"
	"project1/usecases"
	"project1/usecases/domain"
)

type PersonalProfileService interface {
	GetPersonalProfile(ctx context.Context, req domain.PersonalProfile) (res []domain.PersonalProfile, err error)
	CreatePersonalProfile(ctx context.Context, req domain.PersonalProfile) (res domain.PersonalProfile, err error)
	GetAllPersonalProfiles(ctx context.Context) (res []domain.PersonalProfile, err error)
	UpdatePersonalProfile(ctx context.Context, req domain.PersonalProfile) (domain.PersonalProfile, error)
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

func (PersonalProfile) UpdatePersonalProfile(ctx context.Context, req domain.PersonalProfile) (res domain.PersonalProfile, err error) {
	interactor := usecases.PersonalProfileInterface{
		PersonalProfileRepository: repositories.PersonalProfileRepository,
	}

	res, err = interactor.UpdatePersonalProfile(ctx, req)
	return
}
