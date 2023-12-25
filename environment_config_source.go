package yach

import (
	"os"
	"strconv"
)

// EnvironmentConfigSource will retrieve values from environment variables.
type EnvironmentConfigSource struct {
	bindings map[string]string
	autoBind bool
}

func NewEnvironmentConfigSource() *EnvironmentConfigSource {
	return &EnvironmentConfigSource{
		bindings: map[string]string{},
		autoBind: false,
	}
}

// Bind an (abstract) key to an environment variable.
// This method is useful because environment variable names will often not match the names of your actual configuration directives.
func (e *EnvironmentConfigSource) Bind(key string, value string) {
	e.bindings[key] = value
}

// AutoBind enables automatic binding of keys to environment variables.
// If this is enabled and a value is requested for a key that has not been bound to an environment variable using Bind(),
// it will treat the 'key' parameter as an environment variable directly and try to retrieve its value.
func (e *EnvironmentConfigSource) AutoBind() {
	e.autoBind = true
}

func (e *EnvironmentConfigSource) Get(key string) (string, error) {
	// First, look up the key to see if a binding exists
	keyBinding, exists := e.bindings[key]
	if exists {
		value := os.Getenv(keyBinding)
		if value == "" {
			return "", ErrNoValueFound
		}

		return value, nil
	}

	// If a binding does not exist, check if auto binding is enabled
	// If so, look up the raw key itself. If not, return an error.
	if e.autoBind {
		value := os.Getenv(key)
		if value == "" {
			return "", ErrNoValueFound
		}

		return value, nil
	}

	return "", ErrNoBindingFound
}

func (e *EnvironmentConfigSource) GetInt(key string) (int, error) {
	value, err := e.Get(key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(value)
}
