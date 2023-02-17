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

package config

import (
	"github.com/franciscosbf/micro-dwarf/internal/clis"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"time"
)

// PostgresConfig contains all elements to establish a
// connection pool, along with optional parameters
// to control the pool behaviour
type PostgresConfig struct {
	// Connection related

	User     string `name:"POSTGRES_USER_SECRET" required:"yes"`
	Password string `name:"POSTGRES_PASSWORD_SECRET" required:"yes"`
	Host     string `name:"POSTGRES_HOST" required:"yes"`
	Port     uint16 `name:"POSTGRES_PORT"`
	Dbname   string `name:"POSTGRES_DBNAME" required:"yes"`
	SslMode  string `name:"POSTGRES_SSL_MODE" accepts:"disable,allow,prefer,require,verify-ca,verify-full"`

	// Pool configuration

	PoolMaxCons               int32         `name:"POSTGRES_POOL_MAX_CONS"`
	PoolMinCons               int32         `name:"POSTGRES_POOL_MIN_CONS"`
	PoolMaxConnLifetime       time.Duration `name:"POSTGRES_POOL_MAX_CONN_LIFETIME"`
	PoolMaxConnIdleTime       time.Duration `name:"POSTGRES_POOL_MAX_CONN_IDLE_TIME"`
	PoolHealthCheckPeriod     time.Duration `name:"POSTGRES_POOL_HEALTH_CHECK_PERIOD"`
	PoolMaxConnLifetimeJitter time.Duration `name:"POSTGRES_POOL_MAX_CONN_LIFETIME_JITTER"`
}

// New returns a new postgres config
func New(vReader *envvars.VarReader) (template *PostgresConfig, err error) {
	template = &PostgresConfig{}
	err = clis.ReadConfTemplate(vReader, template)

	return
}
