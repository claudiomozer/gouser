package user

import "github.com/claudiomozer/gouser/internal/domain/types"

type Entity struct {
	ID    string      `json:"id"`
	Name  types.Name  `json:"name"`
	Role  types.Role  `json:"role"`
	Email types.Email `json:"email"`
}
