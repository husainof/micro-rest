package user

import (
	"context"
	"micro/internal/user"
	"micro/pkg/logging"
)

type service struct {
	store  Storage
	logger *logging.Logger
}

func NewUserService(store user.Storage, logger *logging.Logger) *service {
	return &service{store, logger}
}

func (s *service) Create(ctx context.Context, dto user.CreateUserDTO) (*User, error) {

	return nil, nil
}
