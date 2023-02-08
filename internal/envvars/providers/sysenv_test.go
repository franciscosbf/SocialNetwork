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

import (
	"os"
	"testing"
)

func TestWithoutErrorWithVariableSet(t *testing.T) {
	_ = os.Setenv("TEST_1", "hi")
	defer func() { _ = os.Unsetenv("TEST_1") }()

	envVars := NewEnvVariables()

	if _, err := envVars.Get("TEST_VAR"); err != nil {
		t.Error("Must not return any error when a variable exist")
	}
}

func TestWithoutErrorWithVariableUnset(t *testing.T) {
	envVars := NewEnvVariables()

	if _, err := envVars.Get("TEST_2"); err != nil {
		t.Error("Must not return any error when a variable doesn't exist")
	}
}

func TestVariableSet(t *testing.T) {
	t.Setenv("TEST_3", "hi")
	defer func() { _ = os.Unsetenv("TEST_3") }()

	envVars := NewEnvVariables()

	if value, err := envVars.Get("TEST_3"); value != "hi" {
		t.Error("Expecting var TEST_3 containing value hi")
	} else if err != nil {
		t.Error("Expecting err to be nil")
	}
}

func TestVariableUnset(t *testing.T) {
	envVars := NewEnvVariables()

	if value, err := envVars.Get("TEST_4"); value != "" {
		t.Error("Expecting returning empty string when a variable doesn't exist")
	} else if err != nil {
		t.Error("Expecting err to be nil")
	}
}
