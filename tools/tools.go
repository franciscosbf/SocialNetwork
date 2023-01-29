//go:build tools
// +build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"       // Database Migrations
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Database Migrations driver

	_ "github.com/jackc/pgx/v5"             // PostgresSQL db driver
	_ "github.com/kyleconroy/sqlc/cmd/sqlc" // Type-Safe SQL generator
)
