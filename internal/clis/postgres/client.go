/*
Copyright 2023 Francisco Simões Braço-Forte

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package postgres

import (
	"context"
	"fmt"
	"github.com/franciscosbf/micro-dwarf/internal/clis"
	"github.com/franciscosbf/micro-dwarf/internal/clis/postgres/config"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/errorw"
	"github.com/franciscosbf/micro-dwarf/internal/utils"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	ErrorCodeQueryCheckFail errorw.ErrorCode = iota
)

// dsnConn returns a dsn containing only connection elements
func dsnConn(varsConf *config.PostgresConfig) (dsn string) {
	dsn = fmt.Sprintf(
		"user=%v password=%v host=%v dbname=%v",
		varsConf.User, varsConf.Password, varsConf.Host, varsConf.Dbname)

	if port := varsConf.Port; port != 0 {
		dsn = fmt.Sprintf("%v port=%v", dsn, port)
	}

	if sslMode := varsConf.SslMode; sslMode != "" {
		dsn = fmt.Sprintf("%v sslmode=%v", dsn, sslMode)
	}

	return
}

// populatePgxDefs sets up all pgxConf parameters if present in varsConf
func populatePgxDefs(varsConf *config.PostgresConfig, pgxConf *pgxpool.Config) {
	utils.SetAny(varsConf.PoolMaxCons, &pgxConf.MaxConns)
	utils.SetAny(varsConf.PoolMinCons, &pgxConf.MinConns)
	utils.SetAny(varsConf.PoolMaxConnLifetime, &pgxConf.MaxConnLifetime)
	utils.SetAny(varsConf.PoolMaxConnIdleTime, &pgxConf.MaxConnIdleTime)
	utils.SetAny(varsConf.PoolHealthCheckPeriod, &pgxConf.HealthCheckPeriod)
	utils.SetAny(varsConf.PoolMaxConnLifetimeJitter, &pgxConf.MaxConnLifetimeJitter)
}

// New creates a new pool and checks db connection
func New(vReader *envvars.VarReader) (*pgxpool.Pool, error) {
	if vReader == nil {
		return nil, errorw.WrapErrorf(
			clis.ErrorCodeMissingConfig, nil, "Postgres config is nil")
	}

	varsConf, err := config.New(vReader)
	if err != nil {
		return nil, errorw.WrapErrorf(
			clis.ErrorCodeVarReader, err, "Couldn't build Postgres variables config")
	}

	dsn := dsnConn(varsConf)

	pgxConf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, errorw.WrapErrorf(
			clis.ErrorCodeClientConfigFail, err, "Invalid Postgres config")
	}

	fmt.Println(pgxConf.ConnConfig.ConnString())

	populatePgxDefs(varsConf, pgxConf)

	pool, err := pgxpool.ConnectConfig(context.Background(), pgxConf)
	if err != nil {
		return nil, errorw.WrapErrorf(
			clis.ErrorCodeConnFail, err, "Couldn't create Postgres pool")
	}

	// Checks if connection is ok
	if err := pool.Ping(context.Background()); err != nil {
		return nil, errorw.WrapErrorf(
			ErrorCodeQueryCheckFail, err, "Couldn't perform query check in Postgres database")
	}

	return pool, err
}
