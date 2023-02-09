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

package dsn

import (
	"github.com/franciscosbf/micro-dwarf/internal/clis/postgres/dsn/vars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars/providers"
	"os"
	"strings"
	"testing"
)

func setVars() {
	_ = vars.ForEachPostgresVar(func(info *vars.PostgresVarInfo) error {
		_ = os.Setenv(info.Name(), "a")

		return nil
	})
}

func unsetVars() {
	_ = vars.ForEachPostgresVar(func(info *vars.PostgresVarInfo) error {
		_ = os.Unsetenv(info.Name())

		return nil
	})
}

func validateDsn(t *testing.T, cond func(info *vars.PostgresVarInfo) bool) {
	envVars := providers.NewEnvVariables()
	conf := envvars.NewConfig(envVars)

	dsn, err := Build(conf)
	if err != nil {
		t.Errorf("Unexpecting error: %v", err)
	}

	// Check if each dsn pair is correctly formatted
	keys := make(map[string]struct{})
	for _, pair := range strings.Split(dsn, " ") {
		sp := strings.Split(pair, "=")

		if len(sp) != 2 {
			t.Errorf("Bad dsn key par: %v. Should be <key>=<value>", pair)
		}

		keys[sp[0]] = struct{}{}
	}

	_ = vars.ForEachPostgresVar(func(info *vars.PostgresVarInfo) error {
		if _, ok := keys[info.Dsn()]; !ok && cond(info) {
			t.Errorf("Missing dsn %v which var name is %v", info.Dsn(), info.Name())
		}

		return nil
	})
}

func TestVars(t *testing.T) {
	setVars()

	tests := map[string]func(t *testing.T){
		"TestWithAllVariables": func(t *testing.T) {
			validateDsn(t, func(info *vars.PostgresVarInfo) bool {
				return true
			})
		},
		"TestWithRequiredVariables": func(t *testing.T) {
			validateDsn(t, func(info *vars.PostgresVarInfo) bool {
				return info.Required()
			})
		},
		"TestWithOptionalVariables": func(t *testing.T) {
			validateDsn(t, func(info *vars.PostgresVarInfo) bool {
				return !info.Required()
			})
		},
	}

	for name, test := range tests {
		t.Run(name, test)
	}

	unsetVars()
}
