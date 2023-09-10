package yach

import (
	"os"
	"strconv"
)

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

func (e *EnvironmentConfigSource) Bind(key string, value string) {
	e.bindings[key] = value
}

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
