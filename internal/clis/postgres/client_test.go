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

	if cli, err := NewPostgresCli(conf); err != nil {
		t.Errorf("Unexpecting error: %v", err)
	} else if cli == nil {
		t.Errorf("Client should not be nil")
	} else {
		cli.Close()
	}
}
