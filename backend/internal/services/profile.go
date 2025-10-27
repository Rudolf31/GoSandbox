package services

import (
	"context"
	"errors"
	customeerrors "interface_lesson/internal/customeErrors"
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
	profilesMap map[int32]dto.ProfileDTO
	pool        *pgxpool.Pool
}

// CreateProfile implements ProfileService.
func (p *profileServiceImpl) CreateProfile(profile dto.NewProfileDTO) (*int32, *customeerrors.Wrapper) {

	q := database.New(p.pool)

	newProfile, err := q.CreateProfile(context.Background(), database.CreateProfileParams{
		Name:     &profile.Name,
		LastName: &profile.LastName,
		Age:      int32(profile.Age),
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
		Name:     *newProfile.Name,
		LastName: *newProfile.LastName,
		Age:      int(newProfile.Age),
	}

	return &DTOProfile, nil

}

// UpdateProfile implements ProfileService.
func (p *profileServiceImpl) UpdateProfile(id int32, profile dto.NewProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper) {
	existing, ok := p.profilesMap[id]
	if !ok {
		return nil, &customeerrors.Wrapper{
			Error:       customeerrors.ErrNotFound,
			ID:          0,
			Description: "We haven't that user..",
		}
	}

	// (без лишних аллокаций)
	existing.Name = profile.Name
	existing.LastName = profile.LastName
	existing.Age = profile.Age

	p.profilesMap[id] = existing // перезапись значения в map

	return &existing, nil
}

func (p *profileServiceImpl) PatchProfile(id int32, profile dto.PatchProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper) {
	existing, ok := p.profilesMap[id]
	if !ok {
		return nil, &customeerrors.Wrapper{
			Error:       customeerrors.ErrNotFound,
			ID:          0,
			Description: "We haven't that user..",
		}
	}

	if profile.Age != nil {
		existing.Age = *profile.Age
	}
	if profile.LastName != nil {
		existing.LastName = *profile.LastName
	}
	if profile.Name != nil {
		existing.Name = *profile.Name
	}

	p.profilesMap[id] = existing

	return &existing, nil
}

func NewProfileService(pool *pgxpool.Pool) ProfileService {
	p := &profileServiceImpl{
		profilesMap: make(map[int32]dto.ProfileDTO),
		pool:        pool,
	}
	return p
}
