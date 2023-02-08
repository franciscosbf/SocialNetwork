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

package envvars

import "github.com/franciscosbf/micro-dwarf/internal/errorw"

const ErrorCodeVarFetch errorw.ErrorCode = 0

// Provider represents the connector
// that fetches variables. If key doesn't
// exist or an error has occurred, it should
// return an empty string
type Provider interface {
	Get(key string) (string, error)
}

// Config wraps the process of getting
// variables with a given Provider
type Config struct {
	provider Provider
}

// NewConfig Creates a new config with a given provider
func NewConfig(provider Provider) *Config {
	return &Config{provider: provider}
}

// Get returns a variable from the provider given its key
func (c *Config) Get(key string) (string, error) {
	value, err := c.provider.Get(key)
	if err != nil {
		return "", errorw.WrapErrorf(ErrorCodeVarFetch, err, "Couldn't get variable %v", key)
	}

	return value, nil
}
