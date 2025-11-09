package services

import (
	"context"
	"errors"
	"interface_lesson/internal/customerrors"
	"interface_lesson/internal/database"
	"interface_lesson/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type AuthService interface {
	Authenticate(username, password string) (*string, *customerrors.Wrapper)
}

type authServiceImpl struct {
	log  *zap.Logger
	pool *pgxpool.Pool
}

func (a *authServiceImpl) Authenticate(username, password string) (*string, *customerrors.Wrapper) {
	q := database.New(a.pool)
	c := context.TODO()

	account, err := q.GetAccountByUsername(c, username)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {

			a.log.Warn(
				"Account doesn't exist",
				zap.String("username", username),
				zap.Error(err),
			)

			return nil, customerrors.NewErrorWrapper(
				customerrors.ErrUnauthorized,
				"",
			)
		}

		a.log.Error(
			"Failed to get account",
			zap.String("username", username),
			zap.Error(err),
		)

		return nil, customerrors.NewErrorWrapper(
			customerrors.ErrServerError,
			"",
		)
	}

	loginInfo, err := q.GetLoginInfo(c, account.ID)
	if err != nil {
		a.log.Error(
			"Failed to get login info",
			zap.String("username", username),
			zap.Error(err),
		)

		return nil, customerrors.NewErrorWrapper(
			customerrors.ErrServerError,
			"",
		)
	}

	hash, _ := utils.HashPassword(password)

	if hash != loginInfo.PasswordHesh {

		a.log.Error(
			"Password incorrect",
			zap.String("username", username),
			zap.Error(err),
		)

		return nil, customerrors.NewErrorWrapper(
			customerrors.ErrUnauthorized,
			"",
		)
	}

	return &account.Username, nil

}

func NewAuthService(pool *pgxpool.Pool, log *zap.Logger) AuthService {
	a := &authServiceImpl{
		pool: pool,
		log:  log,
	}

	return a
}
