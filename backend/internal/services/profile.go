package services

import (
	"context"
	customeerrors "interface_lesson/internal/customeErrors"
	"interface_lesson/internal/database"
	"interface_lesson/internal/models/dto"

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

	// newProfile := dto.ProfileDTO{
	// 	Id:       rand.Int(),
	// 	Name:     profile.Name,
	// 	LastName: profile.LastName,
	// 	Age:      profile.Age,
	// }

	// p.profilesMap[newProfile.Id] = newProfile
}

// DeleteProfile implements ProfileService.
func (p *profileServiceImpl) DeleteProfile(id int32) *customeerrors.Wrapper {
	_, ok := p.profilesMap[id]
	if !ok {
		return &customeerrors.Wrapper{
			Error:       customeerrors.ErrNotFound,
			ID:          0,
			Description: "We haven't that user..",
		}
	}

	delete(p.profilesMap, id)

	return nil
}

// GetProfile implements ProfileService.
func (p *profileServiceImpl) GetProfile(id int32) (*dto.ProfileDTO, *customeerrors.Wrapper) {
	profile, ok := p.profilesMap[id]

	if !ok {
		return nil, &customeerrors.Wrapper{
			Error:       customeerrors.ErrNotFound,
			ID:          0,
			Description: "We haven't that user..",
		}
	}

	return &profile, nil

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
