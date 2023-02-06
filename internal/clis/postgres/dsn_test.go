package postgres

import (
	"fmt"
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars/providers"
	"strings"
	"testing"
)

func setVars(t *testing.T, value string, varPos ...int) {
	for _, varP := range varPos {
		t.Setenv(confVars[varP].varName, value)
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
	setVars(t, "a", 1, 2, 3)

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
	setVars(t, "b", 6, 7, 8)
	setVars(t, "", 9)

	envVars := providers.NewEnvVariables()
	conf := envvars.NewConfig(envVars)

	expectedDsn := buildDummyDsn("b", 6, 7, 8)

	if dsn, err := buildDsn(conf); err != nil {
		t.Errorf("Unexpected error: %v", err)
	} else if dsn != expectedDsn {
		t.Errorf("Expecting dsn: %v. Got: %v", expectedDsn, dsn)
	}
}
