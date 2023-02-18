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

package redis

import (
	"context"
	"fmt"
	"github.com/franciscosbf/micro-dwarf/internal/clis"
	"github.com/franciscosbf/micro-dwarf/internal/clis/redis/config"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/errorw"
	"github.com/go-redis/redis/v8"
)

// Error codes

const (
	ErrorCodeNodeConnFail errorw.ErrorCode = iota
)

// createClusterConf initializes the cluster options, returning it
func createClusterConf(varsConf *config.RedisConfig) *redis.ClusterOptions {
	opts := &redis.ClusterOptions{
		// Fields that receive a value by default, regardless
		// the corresponding variable is defined or not

		Username:           varsConf.Username,
		Password:           varsConf.Password,
		ReadOnly:           varsConf.ReadOnly,
		PoolFIFO:           varsConf.PoolFifo,
		MaxRedirects:       varsConf.MaxRedirects,
		MaxRetries:         varsConf.MaxRetries,
		PoolSize:           varsConf.PoolSize,
		MinIdleConns:       varsConf.MinIdleConnections,
		MinRetryBackoff:    varsConf.MinRetryBackOff,
		MaxRetryBackoff:    varsConf.MaxRetryBackOff,
		DialTimeout:        varsConf.DialTimeout,
		ReadTimeout:        varsConf.ReadTimout,
		WriteTimeout:       varsConf.WriteTimout,
		PoolTimeout:        varsConf.PoolTimeout,
		IdleTimeout:        varsConf.IdleTimeout,
		MaxConnAge:         varsConf.MaxConnAge,
		IdleCheckFrequency: varsConf.IdleCheckFrequency,
	}

	// Add node addresses
	aBucket := varsConf.Addrs.Bucket
	for _, addr := range aBucket {
		formatted := fmt.Sprintf(
			"%v:%v", addr.Host, addr.Port)

		opts.Addrs = append(opts.Addrs, formatted)
	}

	// Select route mode
	switch varsConf.RouteMode {
	case "latency":
		opts.RouteByLatency = true
	case "randomly":
		opts.RouteRandomly = true
	}

	return opts
}

// pingNode checks connection with a given
// node and returns an error if any
func pingNode(ctx context.Context, client *redis.Client) error {
	return client.Ping(ctx).Err()
}

// New creates a new cluster cli and checks connection with all shards
func New(vReader *envvars.VarReader) (*redis.ClusterClient, error) {
	if vReader == nil {
		return nil, errorw.WrapErrorf(
			clis.ErrorCodeMissingReader, nil, "Redis variables reader is nil")
	}

	varsConf, err := config.New(vReader)
	if err != nil {
		return nil, errorw.WrapErrorf(
			clis.ErrorCodeVarReader, err, "Couldn't build Redis variables config")
	}

	cConf := createClusterConf(varsConf)
	cli := redis.NewClusterClient(cConf)

	// Iterates over all nodes to perform a heath-check
	ctx := context.Background()
	if err := cli.ForEachShard(ctx, pingNode); err != nil {
		return nil, errorw.WrapErrorf(
			ErrorCodeNodeConnFail, err, "Failed Redis connection check in a node")
	}

	return cli, nil
}
