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
	"github.com/franciscosbf/micro-dwarf/internal/clis/postgres/dsn/vars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars/providers"
	"os"
	"testing"
)

func unsetVars() {
	_ = vars.ForEachPostgresVar(func(info *vars.PostgresVarInfo) error {
		_ = os.Unsetenv(info.VarName)

		return nil
	})
}

func TestValidConnection(t *testing.T) {
	envProvider := providers.NewEnvVariables()
	conf := envvars.NewConfig(envProvider)

	defer unsetVars()

	if cli, err := NewClient(conf); err != nil {
		t.Errorf("Unexpecting error: %v", err)
	} else if cli == nil {
		t.Errorf("Client should not be nil")
	} else {
		cli.Close()
	}
}
