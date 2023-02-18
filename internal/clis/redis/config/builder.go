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
	"github.com/franciscosbf/micro-dwarf/internal/utils"
	"time"
)

// RedisConfig contains all cluster address nodes
type RedisConfig struct {
	// Connection related

	Addrs    *utils.Addrs `name:"REDIS_ADDRESSES" required:"yes"` // see utils.ParseAddrs
	Username string       `name:"REDIS_USERNAME_SECRET"`
	Password string       `name:"REDIS_PASSWORD_SECRET" required:"yes"`

	// Connection and pool configurations

	RouteMode          string        `name:"REDIS_ROUTE_MODE" accepts:"latency,randomly"`
	ReadOnly           bool          `name:"REDIS_READ_ONY"`
	PoolFifo           bool          `name:"REDIS_POOL_FIFO"`
	MaxRedirects       int           `name:"REDIS_MAX_REDIRECTS"`
	MaxRetries         int           `name:"REDIS_MAX_RETRIES"`
	PoolSize           int           `name:"REDIS_POOL_SIZE"`
	MinIdleConnections int           `name:"REDIS_MIN_IDLE_CONNECTIONS"`
	MinRetryBackOff    time.Duration `name:"REDIS_MIN_RETRY_BACKOFF"`
	MaxRetryBackOff    time.Duration `name:"REDIS_MIN_RETRY_BACKOFF"`
	DialTimeout        time.Duration `name:"REDIS_DIAL_TIMEOUT"`
	ReadTimout         time.Duration `name:"REDIS_READ_TIMEOUT"`
	WriteTimout        time.Duration `name:"REDIS_WRITE_TIMEOUT"`
	PoolTimeout        time.Duration `name:"REDIS_POOL_TIMEOUT"`
	IdleTimeout        time.Duration `name:"REDIS_IDLE_TIMEOUT"`
	MaxConnAge         time.Duration `name:"REDIS_MAX_CONN_AGE"`
	IdleCheckFrequency time.Duration `name:"REDIS_IDLE_CHECK_FREQUENCY"`
}

// New returns a new redis config
func New(vReader *envvars.VarReader) (template *RedisConfig, err error) {
	template = &RedisConfig{}
	err = clis.ReadConfTemplate(vReader, template)

	return
}
