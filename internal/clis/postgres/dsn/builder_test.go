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
		_ = os.Setenv(info.VarName, "a")

		return nil
	})
}

func unsetVars() {
	_ = vars.ForEachPostgresVar(func(info *vars.PostgresVarInfo) error {
		_ = os.Unsetenv(info.VarName)

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
		if _, ok := keys[info.DsnName]; !ok && cond(info) {
			t.Errorf("Missing dsn %v which var name is %v", info.DsnName, info.VarName)
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
				return info.Required
			})
		},
		"TestWithOptionalVariables": func(t *testing.T) {
			validateDsn(t, func(info *vars.PostgresVarInfo) bool {
				return !info.Required
			})
		},
	}

	for name, test := range tests {
		t.Run(name, test)
	}

	unsetVars()
}
