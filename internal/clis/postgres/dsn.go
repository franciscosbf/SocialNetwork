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
	"fmt"
	"github.com/franciscosbf/micro-dwarf/internal/common"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"strings"
)

const (
	ErrorCodeInvalidGetVar common.ErrorCode = 0
)

// DsnConn represents dsn connection values
type DsnConn struct {
	values []string
}

// addString adds a string value to dsn. Empty string is ignored
func (d *DsnConn) add(key string, value string) {
	if value == "" {
		return
	}

	pair := fmt.Sprintf("%v=%v", key, value)
	d.values = append(d.values, pair)
}

// unify builds the dsn string
func (d *DsnConn) unify() string {
	return strings.Join(d.values, " ")
}

// buildDsn returns a valid Postgres connection dsn
func buildDsn(connData *envvars.Config) (string, error) {
	raw := &DsnConn{}

	for _, pair := range confVars {
		value, err := connData.Get(pair.varName)
		if err != nil {
			return "", common.WrapErrorf(
				ErrorCodeInvalidGetVar, err, "Invalid DSN var fetch")
		}

		raw.add(pair.dsnName, value)
	}

	return raw.unify(), nil
}
