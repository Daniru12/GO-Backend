package mysql

import (
	"time"
	"context"
	"project1/database"
	log "project1/logger"
	"project1/usecases/domain"
)

type PersonalProfileRepository struct{}

func (PersonalProfileRepository) GetPersonalProfile(ctx context.Context, req domain.PersonalProfile) (res []domain.PersonalProfile, err error) {
	argsSlice, query := CustomQueryGetPersonalProfile(req.Id)
	res, err = execQuery(ctx, query, argsSlice)
	if err != nil {
		log.ErrorContext(ctx, "GetPersonalProfile.Query", query, argsSlice)
	}
	return res, err
}

func (PersonalProfileRepository) GetAllPersonalProfiles(ctx context.Context) (res []domain.PersonalProfile, err error) {
	query := `SELECT id, name, description, status, create_time, update_time 
	          FROM personal_profiles`

	rows, err := database.Connections.Read.QueryContext(ctx, query)
	if err != nil {
		log.ErrorContext(ctx, "GetAllPersonalProfiles.Query", query, err)
		return nil, err
	}
	defer rows.Close()

	var profiles []domain.PersonalProfile
	for rows.Next() {
		c := domain.PersonalProfile{}
		if err := rows.Scan(
			&c.Id,
			&c.Name,
			&c.Description,
			&c.Status,
			&c.CreateTime,
			&c.UpdateTime,
		); err != nil {
			log.ErrorContext(ctx, "GetAllPersonalProfiles.Scan", err)
			return nil, err
		}
		profiles = append(profiles, c)
	}

	return profiles, nil
}

func (PersonalProfileRepository) CreatePersonalProfile(ctx context.Context, req domain.PersonalProfile) (res domain.PersonalProfile, err error) {
	now := time.Now()
	if req.CreateTime.IsZero() {
		req.CreateTime = now
	}
	if req.UpdateTime.IsZero() {
		req.UpdateTime = now
	}

	query := `INSERT INTO personal_profiles (name, description, status, create_time, update_time) 
			  VALUES (?, ?, ?, ?, ?)`

	result, err := database.Connections.Write.Exec(query, req.Name, req.Description, req.Status, req.CreateTime, req.UpdateTime)
	if err != nil {
		log.ErrorContext(ctx, "CreatePersonalProfile.Exec", query, req)
		return res, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.ErrorContext(ctx, "CreatePersonalProfile.LastInsertId", err)
		return res, err
	}

	res = domain.PersonalProfile{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		CreateTime:  req.CreateTime,
		UpdateTime:  req.UpdateTime,
	}
	return res, nil
}

func (PersonalProfileRepository) UpdatePersonalProfile(ctx context.Context, req domain.PersonalProfile) (res domain.PersonalProfile, err error) {
	query := `UPDATE personal_profiles SET name=?, description=?, status=?, update_time=? WHERE id=?`
	_, err = database.Connections.Write.Exec(query, req.Name, req.Description, req.Status, req.UpdateTime, req.Id)
	if err != nil {
		log.ErrorContext(ctx, "UpdatePersonalProfile.Exec", err)
		return res, err
	}
	res = req
	return res, nil
}

func CustomQueryGetPersonalProfile(req int64) ([]interface{}, string) {
	args := []interface{}{}

	query := `SELECT id, name, description, status, create_time, update_time 
	          FROM personal_profiles`

	if req != 0 {
		query += " WHERE id = ?"
		args = append(args, req)
	}

	return args, query
}

func execQuery(ctx context.Context, query string, args []interface{}) (res []domain.PersonalProfile, err error) {
	payDB := []domain.PersonalProfile{}

	rows, err := database.Connections.Read.Query(query, args...)
	if err != nil {
		log.ErrorContext(ctx, "PersonalProfileRepository.execQuery", err)
		return payDB, err
	}
	defer rows.Close()

	for rows.Next() {
		c := domain.PersonalProfile{}
		err = rows.Scan(
			&c.Id,
			&c.Name,
			&c.Description,
			&c.Status,
			&c.CreateTime,
			&c.UpdateTime,
		)
		if err != nil {
			log.ErrorContext(ctx, "PersonalProfileRepository.execQuery.Scan", err)
			return payDB, err
		}
		payDB = append(payDB, c)
	}

	return payDB, nil
}
