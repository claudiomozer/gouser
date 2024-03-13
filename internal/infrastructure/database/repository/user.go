package repository

import (
	"context"
	"errors"

	"github.com/claudiomozer/gouser/internal/domain/err"
	"github.com/claudiomozer/gouser/internal/domain/user"
	"github.com/claudiomozer/gouser/internal/infrastructure/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (r *UserRepository) Create(ctx context.Context, entity *user.Entity) error {
	query := `
		INSERT INTO users (
			id, name, role, email 
		) VALUES (
			@id, @name, @role, @email
		)
	`
	_, execErr := r.pool.Exec(ctx, query, pgx.NamedArgs{
		"id":    entity.ID,
		"name":  entity.Name,
		"email": entity.Email,
		"role":  entity.Role,
	})

	if execErr != nil {
		pgErr := &pgconn.PgError{}
		if errors.As(execErr, &pgErr) && pgErr.Code == database.UniqueViolationErrCode {
			return err.New(err.ErrUserAlreadyExists, "user already exists")
		}
		return execErr
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, userID string) error {
	query := "DELETE FROM users WHERE id = @id"

	command, execErr := r.pool.Exec(ctx, query, pgx.NamedArgs{
		"id": userID,
	})

	if execErr != nil {
		return execErr
	}

	if command.RowsAffected() == 0 {
		return err.New(err.ErrUserNotExists, "user not exists")
	}

	return nil
}
