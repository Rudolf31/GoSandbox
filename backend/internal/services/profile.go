package services

import (
	customeerrors "interface_lesson/internal/customeErrors"
	"interface_lesson/internal/models/dto"
	"math/rand"
)

type ProfileService interface {
	CreateProfile(profile dto.NewProfileDTO) (*int, *customeerrors.Wrapper)
	GetProfile(id int) (*dto.ProfileDTO, *customeerrors.Wrapper)
	UpdateProfile(id int, profile dto.NewProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper)
	DeleteProfile(id int) *customeerrors.Wrapper
	PatchProfile(id int, profile dto.PatchProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper)
}

type profileServiceImpl struct {
	profilesMap map[int]dto.ProfileDTO
}

// CreateProfile implements ProfileService.
func (p *profileServiceImpl) CreateProfile(profile dto.NewProfileDTO) (*int, *customeerrors.Wrapper) {

	newProfile := dto.ProfileDTO{
		Id:       rand.Int(),
		Name:     profile.Name,
		LastName: profile.LastName,
		Age:      profile.Age,
	}

	p.profilesMap[newProfile.Id] = newProfile

	return &newProfile.Id, nil
}

// DeleteProfile implements ProfileService.
func (p *profileServiceImpl) DeleteProfile(id int) *customeerrors.Wrapper {
	_, ok := p.profilesMap[id]
	if !ok {
		return &customeerrors.Wrapper{
			Error:       customeerrors.ErrNotFound,
			ID:          id,
			Description: "We haven't that user..",
		}
	}

	delete(p.profilesMap, id)

	return nil
}

// GetProfile implements ProfileService.
func (p *profileServiceImpl) GetProfile(id int) (*dto.ProfileDTO, *customeerrors.Wrapper) {
	profile, ok := p.profilesMap[id]

	if !ok {
		return nil, &customeerrors.Wrapper{
			Error:       customeerrors.ErrNotFound,
			ID:          id,
			Description: "We haven't that user..",
		}
	}

	return &profile, nil

}

// UpdateProfile implements ProfileService.
func (p *profileServiceImpl) UpdateProfile(id int, profile dto.NewProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper) {
	existing, ok := p.profilesMap[id]
	if !ok {
		return nil, &customeerrors.Wrapper{
			Error:       customeerrors.ErrNotFound,
			ID:          id,
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

func (p *profileServiceImpl) PatchProfile(id int, profile dto.PatchProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper) {
	existing, ok := p.profilesMap[id]
	if !ok {
		return nil, &customeerrors.Wrapper{
			Error:       customeerrors.ErrNotFound,
			ID:          id,
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

func NewProfileService() ProfileService {
	p := &profileServiceImpl{profilesMap: make(map[int]dto.ProfileDTO)}
	return p
}
