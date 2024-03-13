package user

import (
	"context"

	"github.com/claudiomozer/gouser/internal/domain/err"
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

func (s *Service) Delete(ctx context.Context, userID string) error {
	if _, parseError := uuid.Parse(userID); parseError != nil {
		return err.InvalidField("id")
	}

	return s.repository.Delete(ctx, userID)
}
