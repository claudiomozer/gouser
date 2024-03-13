package user

import (
	"context"

	"github.com/claudiomozer/gouser/internal/domain/types"
)

//go:generate mockgen -source=repository.go -destination=../../../mocks/user_repository.go -package=mocks
type UserRepository interface {
	Create(ctx context.Context, user *Entity) error
	GetUserRole(ctx context.Context, userID string) (types.Role, error)
	UpdateRole(ctx context.Context, userID string, role types.Role) error
	Delete(ctx context.Context, userID string) error
}
