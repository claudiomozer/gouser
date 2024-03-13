package user

import (
	"context"

	"github.com/claudiomozer/gouser/internal/domain/err"
	"github.com/claudiomozer/gouser/internal/domain/types"
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

func (s *Service) UpdateRole(ctx context.Context, request *UpdateRoleRequest) error {
	if validateErr := request.Validate(); validateErr != nil {
		return validateErr
	}
	newRole := types.FromStringRole(request.Role)

	userRole, getErr := s.repository.GetUserRole(ctx, request.UserID)
	if getErr != nil {
		return getErr
	}

	if newRole > userRole {
		return err.New(err.ErrOperationNotAllowed, "user can not receive a higher role")
	}

	if userRole == newRole {
		return nil
	}

	return s.repository.UpdateRole(ctx, request.UserID, userRole.Update(newRole))
}

func (s *Service) Delete(ctx context.Context, userID string) error {
	if _, parseError := uuid.Parse(userID); parseError != nil {
		return err.InvalidField("id")
	}

	return s.repository.Delete(ctx, userID)
}
