package user

import "context"

//go:generate mockgen -source=repository.go -destination=../../../mocks/user_repository.go -package=mocks
type UserRepository interface {
	Create(ctx context.Context, user *Entity) error
}
