package providers

import (
	"os"
	"testing"
)

func TestWithoutErrorWithVariableSet(t *testing.T) {
	_ = os.Setenv("TEST_1", "hi")

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
	_ = os.Setenv("TEST_3", "hi")

	envVars := NewEnvVariables()

	if value, _ := envVars.Get("TEST_3"); value != "hi" {
		t.Error("Expecting var TEST_3 containing value hi")
	}
}

func TestVariableUnset(t *testing.T) {
	envVars := NewEnvVariables()

	if value, _ := envVars.Get("TEST_4"); value != "" {
		t.Error("Expecting returning empty string when a variable doesn't exist")
	}
}