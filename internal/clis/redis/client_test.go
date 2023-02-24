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
	"github.com/franciscosbf/micro-dwarf/internal/clis"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars/providers"
	"github.com/franciscosbf/micro-dwarf/internal/errorw"
	"github.com/redis/go-redis/v9"
	"os"
	"strings"
	"testing"
)

func unsetVars() {
	for _, v := range os.Environ() {
		key := strings.Split(v, "=")[0]
		if strings.HasPrefix(key, "REDIS_") {
			_ = os.Unsetenv(key)
		}
	}
}

func setVar(key, value string) {
	_ = os.Setenv(key, value)
}

func checkConn(t *testing.T) {
	defer unsetVars()

	envProvider := providers.NewEnvVariables()
	reader := envvars.New(envProvider)

	cli, err := New(reader)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if cli == nil {
		t.Errorf("Unexpect nil client")
	}
}

func TestValidConnection(t *testing.T) {
	checkConn(t)
}

func TestValidSecureConnection(t *testing.T) {
	setVar("REDIS_TLS", "true")

	checkConn(t)
}

func TestInvalidConnection(t *testing.T) {
	checkErrorCode := func(t *testing.T, cli *redis.ClusterClient, err error, code errorw.ErrorCode, errorName string) {
		errw, ok := err.(*errorw.Wrapper)
		if !ok {
			t.Errorf("Expecting errorw.Wrapper, got %v", err)
			return
		}

		if errw.Code() != code {
			t.Errorf("Expecting error code %v, got: %v", errorName, errw.String())
		}

		if cli != nil {
			t.Errorf("Client should be nil, got %v", cli)
		}
	}

	testBattery := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "TestMissingVarsReader",
			test: func(t *testing.T) {
				cli, err := New(nil)
				checkErrorCode(t, cli, err, clis.ErrorCodeMissingReader, "ErrorCodeMissingReader")
			},
		},
		{
			name: "TestMissingVarsReader",
			test: func(t *testing.T) {
				defer unsetVars()

				envProvider := providers.NewEnvVariables()
				reader := envvars.New(envProvider)

				cli, err := New(reader)
				checkErrorCode(t, cli, err, clis.ErrorCodeVarReader, "ErrorCodeVarReader")
			},
		},
		{
			name: "TestInvalidTls",
			test: func(t *testing.T) {
				defer unsetVars()

				setVar("REDIS_ADDRS", "127.255.254.123:1234")
				setVar("REDIS_TLS", "true")

				envProvider := providers.NewEnvVariables()
				reader := envvars.New(envProvider)

				cli, err := New(reader)
				checkErrorCode(t, cli, err, clis.ErrorCodeClientConfigFail, "ErrorCodeClientConfigFail")
			},
		},
		{
			name: "TestClusterConnFailure",
			test: func(t *testing.T) {
				defer unsetVars()

				setVar("REDIS_ADDRS", "127.255.254.123:1234")

				envProvider := providers.NewEnvVariables()
				reader := envvars.New(envProvider)

				cli, err := New(reader)
				checkErrorCode(t, cli, err, ErrorCodeNodeConnFail, "ErrorCodeNodeConnFail")
			},
		},
	}

	for _, pair := range testBattery {
		t.Run(pair.name, pair.test)
	}
}
