package user

import (
	"strings"

	"github.com/claudiomozer/gouser/internal/domain/err"
	"github.com/claudiomozer/gouser/internal/domain/types"
	"github.com/google/uuid"
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

type UpdateRoleRequest struct {
	UserID string           `json:"id"`
	Role   types.StringRole `json:"role"`
}

func (r *UpdateRoleRequest) Validate() error {
	switch {
	case strings.TrimSpace(r.UserID) == "":
		return err.MissingRequiredField("id")
	case strings.TrimSpace(r.Role.String()) == "":
		return err.MissingRequiredField("role")
	}

	if _, validateErr := uuid.Parse(r.UserID); validateErr != nil {
		return err.InvalidField("id")
	}

	if validateErr := r.Role.Validate(); validateErr != nil {
		return validateErr
	}

	return nil
}

type GetUsersRequest struct {
	UserID *string           `json:"id"`
	Name   *types.Name       `json:"name"`
	Role   *types.StringRole `json:"role"`
	Email  *types.Email      `json:"email"`
}

func (r *GetUsersRequest) Validate() error {
	switch {
	case r.UserID != nil && strings.TrimSpace(*r.UserID) == "":
		return err.MissingRequiredField("id")
	case r.Name != nil && strings.TrimSpace(r.Name.String()) == "":
		return err.InvalidField("name")
	case r.Email != nil && strings.TrimSpace(r.Email.String()) == "":
		return err.InvalidField("email")
	case r.Role != nil && strings.TrimSpace(r.Role.String()) == "":
		return err.InvalidField("role")
	}

	if r.Name != nil {
		if validateErr := r.Name.Validate(); validateErr != nil {
			return validateErr
		}
	}

	if r.Email != nil {
		if validateErr := r.Email.Validate(); validateErr != nil {
			return validateErr
		}
	}

	if r.Role != nil {
		if validateErr := r.Role.Validate(); validateErr != nil {
			return validateErr
		}
	}

	return nil
}
