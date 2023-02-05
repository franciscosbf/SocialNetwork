package postgres

import (
	"context"
	"github.com/franciscosbf/micro-dwarf/internal/common"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Error codes
const (
	ErrorCodeConfigFail common.ErrorCode = iota
	ErrorCodeConnBuild
	ErrorCodeConnTry
)

// NewPostgresCli creates a new pool and checks db connection.
func NewPostgresCli(connData *Connection) (*pgxpool.Pool, error) {
	dsn := BuildDsn(connData)

	conf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, common.WrapErrorf(
			ErrorCodeConfigFail, err, "Invalid Postgres config")
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		return nil, common.WrapErrorf(
			ErrorCodeConnBuild, err, "Couldn't create Postgres pool")
	}

	// Checks if connection is ok
	if err := pool.Ping(context.Background()); err != nil {
		return nil, common.WrapErrorf(
			ErrorCodeConnTry, err, "Couldn't connect to Postgres database %v", connData.Dbname)
	}

	return pool, err
}
