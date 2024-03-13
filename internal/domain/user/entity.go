package user

import "github.com/claudiomozer/gouser/internal/domain/types"

type Entity struct {
	ID    string
	Name  types.Name
	Role  types.Role
	Email types.Email
}
