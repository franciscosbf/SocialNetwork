package postgres

import (
	"fmt"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars/providers"
	"os"
	"strings"
	"testing"
)

func setVars(value string, varPos ...int) {
	for _, varP := range varPos {
		_ = os.Setenv(confVars[varP].varName, value)
	}
}

func buildDummyDsn(value string, varPos ...int) string {
	var dsnValues []string
	for _, varP := range varPos {
		pair := fmt.Sprintf("%v=%v", confVars[varP].dsnName, value)
		dsnValues = append(dsnValues, pair)
	}

	return strings.Join(dsnValues, " ")
}

func unsetVars(varPos ...int) {
	for _, varP := range varPos {
		_ = os.Unsetenv(confVars[varP].varName)
	}
}

func TestEmptyDsn(t *testing.T) {
	envVars := providers.NewEnvVariables()
	conf := envvars.NewConfig(envVars)

	if dsn, err := buildDsn(conf); err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if dsn != "" {
		t.Errorf("Expecting empty dsn. Got: %v", dsn)
	}
}

func TestDsnWithSomePairs(t *testing.T) {
	setVars("a", 1, 2, 3)
	defer unsetVars(1, 2, 3)

	envVars := providers.NewEnvVariables()
	conf := envvars.NewConfig(envVars)

	expectedDsn := buildDummyDsn("a", 1, 2, 3)

	if dsn, err := buildDsn(conf); err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if dsn != expectedDsn {
		t.Errorf("Expecting dsn: %v. Got: %v", expectedDsn, dsn)
	}

}

func TestDsnWithEmptyVar(t *testing.T) {
	setVars("b", 6, 7, 8)
	setVars("", 9)
	defer unsetVars(6, 7, 8, 9)

	envVars := providers.NewEnvVariables()
	conf := envvars.NewConfig(envVars)

	expectedDsn := buildDummyDsn("b", 6, 7, 8)

	if dsn, err := buildDsn(conf); err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if dsn != expectedDsn {
		t.Errorf("Expecting dsn: %v. Got: %v", expectedDsn, dsn)
	}
}
