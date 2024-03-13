package app

import (
	"context"

	"github.com/claudiomozer/gouser/internal/domain/user"
	"github.com/claudiomozer/gouser/internal/infrastructure/database"
	"github.com/claudiomozer/gouser/internal/infrastructure/database/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	pgPool      *pgxpool.Pool
	UserService *user.Service
}

func LoadContainer() *Container {
	LoadVars()

	pool := database.ConnectToDatabase(context.Background(), &database.ConnectionConfig{
		UserName:       ENV.DBUser,
		Password:       ENV.DBPassword,
		Host:           ENV.DBHost,
		Port:           ENV.DBPort,
		DBName:         ENV.DBName,
		MaxConnections: ENV.MaxPoolConnections,
	})

	userRepository := repository.NewUserRepository(pool)
	userService := user.NewService(userRepository)

	return &Container{
		pgPool:      pool,
		UserService: userService,
	}
}

func (c *Container) CloseContainer() {
	c.pgPool.Close()
}
