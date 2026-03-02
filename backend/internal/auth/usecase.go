package auth

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	"backend/pkg/hash"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type UseCase struct {
	logger *slog.Logger
	repo   *postgres.Repository
}

func NewUseCase(r *postgres.Repository, l *slog.Logger) *UseCase {
	return &UseCase{
		logger: l,
		repo:   r,
	}
}

func (u *UseCase) Login(ctx context.Context, req *LoginRequest) (domain.UUID, error) {
	user, err := domain.NewUser(req.Username, req.Password)
	if errors.Is(err, domain.ValidationErr) {
		return "", domain.ValidationErr
	} else if err != nil {
		return "", fmt.Errorf("domain user creation error: %w", err)
	}

	dbUser, err := u.repo.GetUser(ctx, user.Username)
	if errors.Is(err, pgx.ErrNoRows) {
		return "", domain.BadCredentialsErr
	} else if err != nil {
		return "", fmt.Errorf("user request error: %v", err)
	}

	if !hash.VerifyPassword(dbUser.Password, req.Password, dbUser.Salt) {
		return "", domain.BadCredentialsErr
	}

	return dbUser.ID, nil
}

func (u *UseCase) Registration(ctx context.Context, req *RegistrationRequest) (domain.UUID, error) {
	user, err := domain.NewUser(req.Username, req.Password)
	if errors.Is(err, domain.ValidationErr) {
		return "", domain.ValidationErr
	} else if err != nil {
		return "", fmt.Errorf("domain user creation error: %w", err)
	}

	isUserExists, err := u.repo.IsUserExists(ctx, *user)
	if err != nil {
		return "", fmt.Errorf("user request error: %v", err)
	}

	if isUserExists {
		return "", domain.ConflictErr
	}

	user.Password, user.Salt, err = hash.HashPassword(req.Password)
	if err != nil {
		return "", fmt.Errorf("hash password error: %v", err)
	}

	uuid, err := u.repo.Registration(ctx, *user)
	if err != nil {
		return "", fmt.Errorf("registration request error: %v", err)
	}

	return uuid, nil
}
