package domain

import (
	"context"
	"project1/util"
)

type PersonalProfileRequest struct {
	PersonalProfile PersonalProfile
}

type CreatePersonalProfileRequest struct {
	PersonalProfile PersonalProfile
}

type PersonalProfile struct {
	Id          int64     `json:"id"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Status      *string    `json:"status"`
	CreateTime  utils.CustomTime `json:"create_time"`
	UpdateTime  utils.CustomTime `json:"update_time"`
}

// Repository
type PersonalProfileRepository interface {
	GetPersonalProfile(ctx context.Context, req PersonalProfile) (res []PersonalProfile, err error)
	GetAllPersonalProfiles(ctx context.Context) (res []PersonalProfile, err error)
	CreatePersonalProfile(ctx context.Context, req PersonalProfile) (res PersonalProfile, err error)
	UpdatePersonalProfile(ctx context.Context, req PersonalProfile) (res PersonalProfile, err error)
}
