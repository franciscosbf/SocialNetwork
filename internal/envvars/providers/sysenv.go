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
