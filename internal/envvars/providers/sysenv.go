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

package providers

import "os"

type EnvVariables struct {
	cache map[string]string
}

// NewEnvVariables creates a new env variables provider
func NewEnvVariables() *EnvVariables {
	varsCache := make(map[string]string)

	return &EnvVariables{cache: varsCache}
}

func (ev *EnvVariables) Get(key string) (string, error) {
	if value, ok := ev.cache[key]; ok {
		return value, nil
	} else {
		value = os.Getenv(key)
		ev.cache[key] = value
		return value, nil
	}
}
