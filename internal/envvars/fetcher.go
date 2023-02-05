package envvars

import "github.com/franciscosbf/micro-dwarf/internal/common"

const ErrorCodeVarFetch common.ErrorCode = 0

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
		return "", common.WrapErrorf(ErrorCodeVarFetch, err, "Couldn't get variable %v", key)
	}

	return value, nil
}
