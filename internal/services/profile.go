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

	return &newProfile.Id, &customeerrors.Wrapper{}
}

// DeleteProfile implements ProfileService.
func (p *profileServiceImpl) DeleteProfile(id int) *customeerrors.Wrapper {
	panic("unimplemented")
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
	panic("unimplemented")
}

func NewProfileService() ProfileService {
	p := &profileServiceImpl{profilesMap: make(map[int]dto.ProfileDTO)}
	return p
}
