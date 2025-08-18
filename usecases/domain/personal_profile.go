package domain

import (
	"context"
	"time"
)

type PersonalProfileRequest struct {
	PersonalProfile PersonalProfile
}

type CreatePersonalProfileRequest struct {
	PersonalProfile PersonalProfile
}

type PersonalProfile struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

// Repository
type PersonalProfileRepository interface {
	GetPersonalProfile(ctx context.Context, req PersonalProfile) (res []PersonalProfile, err error)
	GetAllPersonalProfiles(ctx context.Context) (res []PersonalProfile, err error)
	CreatePersonalProfile(ctx context.Context, req PersonalProfile) (res PersonalProfile, err error)
	UpdatePersonalProfile(ctx context.Context, req PersonalProfile) (res PersonalProfile, err error)
}
