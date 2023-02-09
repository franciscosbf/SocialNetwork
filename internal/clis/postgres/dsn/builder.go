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
	"fmt"
	"github.com/franciscosbf/micro-dwarf/internal/clis"
	"github.com/franciscosbf/micro-dwarf/internal/clis/postgres/dsn/vars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/errorw"
	"strings"
)

// dsnConn represents dsn connection values
type dsnConn struct {
	values []string
}

// addString adds a string value to dsn.
func (d *dsnConn) add(key string, value string) {
	pair := fmt.Sprintf("%v=%v", key, value)
	d.values = append(d.values, pair)
}

// unify builds the dsn string
func (d *dsnConn) unify() string {
	return strings.Join(d.values, " ")
}

// newBuilder returns a new dsn constructor
func newBuilder() *dsnConn {
	return &dsnConn{}
}

// Build returns a valid Postgres connection dsn
func Build(confReader *envvars.Config) (string, error) {
	raw := newBuilder()

	err := vars.ForEachPostgresVar(func(info *vars.PostgresVarInfo) error {
		name := info.VarName
		value, err := confReader.Get(name)
		if err != nil {
			return errorw.WrapErrorf(
				clis.ErrorCodeInvalidGetVar, err, "Invalid Postgres DSN var fetch")
		}

		if value == "" {
			if info.Required {
				return errorw.WrapErrorf(
					clis.ErrorCodeMissingVar, nil, "Missing Postgres required variable %v", name)
			}
		} else {
			raw.add(info.DsnName, value)
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return raw.unify(), nil
}
