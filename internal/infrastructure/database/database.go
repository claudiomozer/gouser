package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const UniqueViolationErrCode = "23505"

type ConnectionConfig struct {
	UserName       string
	Password       string
	Host           string
	DBName         string
	Port           int
	MaxConnections int
}

func ConnectToDatabase(ctx context.Context, config *ConnectionConfig) *pgxpool.Pool {
	pgxConfig, parseConfigErr := pgxpool.ParseConfig(getConnectionURL(config))
	if parseConfigErr != nil {
		panic(parseConfigErr)
	}
	pgxConfig.MaxConns = int32(config.MaxConnections)

	pool, connectionErr := pgxpool.NewWithConfig(ctx, pgxConfig)
	if connectionErr != nil {
		panic(connectionErr)
	}

	return pool
}

func getConnectionURL(config *ConnectionConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
}
