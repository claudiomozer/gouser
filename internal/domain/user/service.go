package user

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repository UserRepository
}

func NewService(repository UserRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, request *CreateRequest) error {
	if validateErr := request.Validate(); validateErr != nil {
		return validateErr
	}

	user := request.ToEntity()
	user.ID = uuid.NewString()

	return s.repository.Create(ctx, user)
}
