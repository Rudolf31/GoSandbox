package services

import (
	"context"
	"errors"
	customeerrors "interface_lesson/internal/customerrors"
	"interface_lesson/internal/database"
	"interface_lesson/internal/models/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileService interface {
	CreateProfile(profile dto.NewProfileDTO) (*int32, *customeerrors.Wrapper)
	GetProfile(id int32) (*dto.ProfileDTO, *customeerrors.Wrapper)
	UpdateProfile(id int32, profile dto.NewProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper)
	DeleteProfile(id int32) *customeerrors.Wrapper
	PatchProfile(id int32, profile dto.PatchProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper)
}

type profileServiceImpl struct {
	pool *pgxpool.Pool
}

// CreateProfile implements ProfileService.
func (p *profileServiceImpl) CreateProfile(profile dto.NewProfileDTO) (*int32, *customeerrors.Wrapper) {

	q := database.New(p.pool)

	newProfile, err := q.CreateProfile(context.Background(), database.CreateProfileParams{
		Name:     profile.Name,
		LastName: profile.LastName,
		Age:      int16(profile.Age),
	})
	if err != nil {
		return nil, &customeerrors.Wrapper{
			Error:       customeerrors.ErrServerError,
			Description: err.Error(),
			ID:          0,
		}
	}

	return &newProfile.ID, nil
}

// DeleteProfile implements ProfileService.
func (p *profileServiceImpl) DeleteProfile(id int32) *customeerrors.Wrapper {

	q := database.New(p.pool)

	_, err := q.DeleteProfile(context.Background(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &customeerrors.Wrapper{
				Error:       customeerrors.ErrNotFound,
				Description: "User not found",
				ID:          0,
			}
		}
		return &customeerrors.Wrapper{
			Error: customeerrors.ErrServerError,
			ID:    0,
		}
	}

	return nil
}

// GetProfile implements ProfileService.
func (p *profileServiceImpl) GetProfile(id int32) (*dto.ProfileDTO, *customeerrors.Wrapper) {

	q := database.New(p.pool)

	newProfile, err := q.GetProfile(context.Background(), id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &customeerrors.Wrapper{
				Error:       customeerrors.ErrNotFound,
				Description: "User not found",
				ID:          0,
			}
		}
		return nil, &customeerrors.Wrapper{
			Error: customeerrors.ErrServerError,
			ID:    0,
		}
	}

	DTOProfile := dto.ProfileDTO{
		Id:       newProfile.ID,
		Name:     newProfile.Name,
		LastName: newProfile.LastName,
		Age:      newProfile.Age,
	}

	return &DTOProfile, nil

}

// UpdateProfile implements ProfileService.
func (p *profileServiceImpl) UpdateProfile(id int32, profile dto.NewProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper) {

	q := database.New(p.pool)

	newProfile, err := q.UpdateProfile(context.Background(), database.UpdateProfileParams{
		ID:       id,
		Name:     profile.Name,
		LastName: profile.LastName,
		Age:      int16(profile.Age),
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &customeerrors.Wrapper{
				Error:       customeerrors.ErrNotFound,
				Description: "User not found",
				ID:          0,
			}
		}
		return nil, &customeerrors.Wrapper{
			Error: customeerrors.ErrServerError,
			ID:    0,
		}
	}

	DTOProfile := dto.ProfileDTO{
		Id:       newProfile.ID,
		Name:     newProfile.Name,
		LastName: newProfile.LastName,
		Age:      int16(newProfile.Age),
	}

	return &DTOProfile, nil
}

func (p *profileServiceImpl) PatchProfile(id int32, profile dto.PatchProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper) {

	q := database.New(p.pool)

	newProfile, err := q.PatchProfile(context.Background(), database.PatchProfileParams{
		ID:       id,
		Name:     profile.Name,
		LastName: profile.LastName,
		Age:      profile.Age,
	})

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, &customeerrors.Wrapper{
				Error:       customeerrors.ErrNotFound,
				Description: "User not found",
				ID:          0,
			}
		}
		return nil, &customeerrors.Wrapper{
			Error: customeerrors.ErrServerError,
			ID:    0,
		}
	}

	DTOProfile := dto.ProfileDTO{
		Id:       newProfile.ID,
		Name:     newProfile.Name,
		LastName: newProfile.LastName,
		Age:      int16(newProfile.Age),
	}

	return &DTOProfile, nil
}

func NewProfileService(pool *pgxpool.Pool) ProfileService {
	p := &profileServiceImpl{
		pool: pool,
	}
	return p
}
