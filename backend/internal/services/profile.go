package services

import (
	"context"
	"errors"
	customeerrors "interface_lesson/internal/customerrors"
	"interface_lesson/internal/database"
	"interface_lesson/internal/models/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type ProfileService interface {
	CreateProfile(profile dto.NewProfileDTO) (*int32, *customeerrors.Wrapper)
	GetProfile(id int32) (*dto.ProfileDTO, *customeerrors.Wrapper)
	UpdateProfile(id int32, profile dto.NewProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper)
	DeleteProfile(id int32) *customeerrors.Wrapper
	PatchProfile(id int32, profile dto.PatchProfileDTO) (*dto.ProfileDTO, *customeerrors.Wrapper)
}

type profileServiceImpl struct {
	log  *zap.Logger
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
		p.log.Error("Failed to create new profile")
		return nil, &customeerrors.Wrapper{
			Error:       customeerrors.ErrServerError,
			Description: err.Error(),
		}
	}

	p.log.Info(
		"New profile created",
		zap.Int32("id", newProfile.ID),
	)

	return &newProfile.ID, nil
}

// DeleteProfile implements ProfileService.
func (p *profileServiceImpl) DeleteProfile(id int32) *customeerrors.Wrapper {

	q := database.New(p.pool)

	_, err := q.DeleteProfile(context.Background(), id)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {

			p.log.Warn(
				"Profile doesn't exist",
				zap.Int32("id", id),
				zap.Error(err),
			)

			return &customeerrors.Wrapper{
				Error:       customeerrors.ErrNotFound,
				Description: "User not found",
			}
		}

		p.log.Error(
			"Failed to delete profile",
			zap.Int32("id", id),
			zap.Error(err),
		)

		return &customeerrors.Wrapper{
			Error: customeerrors.ErrServerError,
		}
	}

	p.log.Info(
		"Profile deleted",
		zap.Int32("id", id),
	)

	return nil
}

// GetProfile implements ProfileService.
func (p *profileServiceImpl) GetProfile(id int32) (*dto.ProfileDTO, *customeerrors.Wrapper) {

	q := database.New(p.pool)

	newProfile, err := q.GetProfile(context.Background(), id)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {

			p.log.Warn(
				"Profile doesn't exist",
				zap.Int32("id", id),
				zap.Error(err),
			)

			return nil, &customeerrors.Wrapper{
				Error:       customeerrors.ErrNotFound,
				Description: "User not found",
			}
		}

		p.log.Error(
			"Failed to get profile",
			zap.Int32("id", id),
			zap.Error(err),
		)

		return nil, &customeerrors.Wrapper{
			Error: customeerrors.ErrServerError,
		}
	}

	DTOProfile := dto.ProfileDTO{
		Id:       newProfile.ID,
		Name:     newProfile.Name,
		LastName: newProfile.LastName,
		Age:      newProfile.Age,
	}

	p.log.Info(
		"Profile retrieved successfully",
		zap.Int32("id", id),
	)

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

			p.log.Warn(
				"Profile doesn't exist",
				zap.Int32("id", id),
				zap.Error(err),
			)

			return nil, &customeerrors.Wrapper{
				Error:       customeerrors.ErrNotFound,
				Description: "User not found",
			}
		}

		p.log.Error(
			"Failed to update profile",
			zap.Int32("id", id),
			zap.Error(err),
		)

		return nil, &customeerrors.Wrapper{
			Error: customeerrors.ErrServerError,
		}
	}

	DTOProfile := dto.ProfileDTO{
		Id:       newProfile.ID,
		Name:     newProfile.Name,
		LastName: newProfile.LastName,
		Age:      int16(newProfile.Age),
	}

	p.log.Info(
		"Profile updated successfully",
		zap.Int32("id", id),
	)

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

			p.log.Warn(
				"Profile doesn't exist",
				zap.Int32("id", id),
				zap.Error(err),
			)

			return nil, &customeerrors.Wrapper{
				Error:       customeerrors.ErrNotFound,
				Description: "User not found",
			}
		}

		p.log.Error(
			"Failed to update profile",
			zap.Int32("id", id),
			zap.Error(err),
		)

		return nil, &customeerrors.Wrapper{
			Error: customeerrors.ErrServerError,
		}
	}

	DTOProfile := dto.ProfileDTO{
		Id:       newProfile.ID,
		Name:     newProfile.Name,
		LastName: newProfile.LastName,
		Age:      int16(newProfile.Age),
	}

	p.log.Info(
		"Profile updated successfully",
		zap.Int32("id", id),
	)

	return &DTOProfile, nil
}

func NewProfileService(pool *pgxpool.Pool, log *zap.Logger) ProfileService {
	p := &profileServiceImpl{
		pool: pool,
		log:  log,
	}
	return p
}
