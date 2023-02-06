package postgres

import (
	"github.com/franciscosbf/micro-dwarf/internal/envvars"
	"github.com/franciscosbf/micro-dwarf/internal/envvars/providers"
	"testing"
)

func TestValidConnection(t *testing.T) {
	envProvider := providers.NewEnvVariables()
	conf := envvars.NewConfig(envProvider)

	if cli, err := NewPostgresCli(conf); err != nil {
		t.Errorf("Unexpecting error: %v", err)
	} else if cli == nil {
		t.Errorf("Client should not be nil")
	} else {
		cli.Close()
	}
}
