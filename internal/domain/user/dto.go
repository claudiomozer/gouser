package user

import (
	"strings"

	"github.com/claudiomozer/gouser/internal/domain/err"
	"github.com/claudiomozer/gouser/internal/domain/types"
)

type CreateRequest struct {
	Name  types.Name       `json:"name"`
	Role  types.StringRole `json:"role"`
	Email types.Email      `json:"email"`
}

func (r *CreateRequest) Validate() error {
	switch {
	case strings.TrimSpace(r.Name.String()) == "":
		return err.MissingRequiredField("name")
	case strings.TrimSpace(r.Email.String()) == "":
		return err.MissingRequiredField("email")
	case strings.TrimSpace(r.Role.String()) == "":
		return err.MissingRequiredField("role")
	}

	if validateErr := r.Name.Validate(); validateErr != nil {
		return validateErr
	}

	if validateErr := r.Email.Validate(); validateErr != nil {
		return validateErr
	}

	if validateErr := r.Role.Validate(); validateErr != nil {
		return validateErr
	}

	return nil
}

func (r *CreateRequest) ToEntity() *Entity {
	return &Entity{
		Name:  r.Name,
		Email: r.Email,
		Role:  types.FromStringRole(r.Role),
	}
}
