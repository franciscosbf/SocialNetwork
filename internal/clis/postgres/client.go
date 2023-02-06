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
	"github.com/franciscosbf/micro-dwarf/internal/clis"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/errorw"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Error codes
const (
	ErrorCodeDsnFail errorw.Code = iota
	ErrorCodeConfigFail
	ErrorCodeConnFail
	ErrorCodeQueryCheckFail
)

// NewPostgresCli creates a new pool and checks db connection.
func NewPostgresCli(connData *envvars.Config) (*pgxpool.Pool, error) {
	if connData == nil {
		return nil, errorw.WrapErrorf(
			clis.ErrorCodeMissingConfig, nil, "Postgres config is nil")
	}

	dsn, err := buildDsn(connData)
	if err != nil {
		return nil, errorw.WrapErrorf(ErrorCodeDsnFail, err, "Couldn't build dsn")
	}

	conf, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, errorw.WrapErrorf(
			ErrorCodeConfigFail, err, "Invalid Postgres config")
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		return nil, errorw.WrapErrorf(
			ErrorCodeConnFail, err, "Couldn't create Postgres pool")
	}

	// Checks if connection is ok
	if err := pool.Ping(context.Background()); err != nil {
		return nil, errorw.WrapErrorf(
			ErrorCodeQueryCheckFail, err, "Couldn't perform a connection query check in Postgres database")
	}

	return pool, err
}
