package yach

import (
	"errors"
	"os"
	"strconv"
	"testing"
)

func TestEnvironmentConfigSource_Get(t *testing.T) {
	e := NewEnvironmentConfigSource()

	key := "key"
	value := "value"
	env := "TEST_KEY"

	err := os.Setenv(env, value)
	if err != nil {
		t.Fatalf("error occurred during call to os.Setenv(), error: %v", err)
	}

	defer func() {
		err = os.Unsetenv(env)
		if err != nil {
			t.Fatalf("error occurred during call to os.Unsetenv(), error: %v", err)
		}
	}()

	e.Bind(key, env)

	actual, err := e.Get(key)
	if err != nil {
		t.Fatalf("expected value %s for key %s, got error %v", value, key, err)
	}

	if actual != value {
		t.Fatalf("expected value %s for key %s, got value %s", value, key, actual)
	}
}

func TestEnvironmentConfigSource_GetErrNoValueFound(t *testing.T) {
	e := NewEnvironmentConfigSource()

	key := "key"
	env := "TEST_KEY"

	e.Bind(key, env)

	actual, err := e.Get(key)
	if err == nil {
		t.Fatalf("expected ErrNoValueFound, got value %s", actual)
	}

	if !errors.Is(err, ErrNoValueFound) {
		t.Fatalf("expected ErrNoValueFound, got error %v", err)
	}
}

func TestEnvironmentConfigSource_GetErrNoBindingFound(t *testing.T) {
	e := NewEnvironmentConfigSource()

	actual, err := e.Get("unknown")
	if err == nil {
		t.Fatalf("expected ErrNoBindingFound, got value %s", actual)
	}

	if !errors.Is(err, ErrNoBindingFound) {
		t.Fatalf("expected ErrNoBindingFound, got error %v", err)
	}
}

func TestEnvironmentConfigSource_GetAutoBind(t *testing.T) {
	e := NewEnvironmentConfigSource()
	e.AutoBind()

	env := "TEST_KEY"
	value := "value"

	err := os.Setenv(env, value)
	if err != nil {
		t.Fatalf("error occurred during call to os.Setenv(), error: %v", err)
	}

	defer func() {
		err = os.Unsetenv(env)
		if err != nil {
			t.Fatalf("error occurred during call to os.Unsetenv(), error: %v", err)
		}
	}()

	actual, err := e.Get(env)
	if err != nil {
		t.Fatalf("expected value %s for key %s, got error %v", value, env, err)
	}

	if actual != value {
		t.Fatalf("expected value %s for key %s, got value %s", value, env, actual)
	}
}

func TestEnvironmentConfigSource_GetAutoBindErrNoValueFound(t *testing.T) {
	e := NewEnvironmentConfigSource()
	e.AutoBind()

	actual, err := e.Get("TEST_KEY")
	if err == nil {
		t.Fatalf("expected ErrNoValueFound, got value %s", actual)
	}

	if !errors.Is(err, ErrNoValueFound) {
		t.Fatalf("expected ErrNoValueFound, got error %v", err)
	}
}

func TestEnvironmentConfigSource_GetInt(t *testing.T) {
	e := NewEnvironmentConfigSource()

	key := "key"
	value := 1
	env := "TEST_KEY"

	err := os.Setenv(env, strconv.Itoa(value))
	if err != nil {
		t.Fatalf("error occurred during call to os.Setenv(), error: %v", err)
	}

	defer func() {
		err = os.Unsetenv(env)
		if err != nil {
			t.Fatalf("error occurred during call to os.Unsetenv(), error: %v", err)
		}
	}()

	e.Bind(key, env)

	actual, err := e.GetInt(key)
	if err != nil {
		t.Fatalf("expected value %d for key %s, got error %v", value, key, err)
	}

	if actual != value {
		t.Fatalf("expected value %d for key %s, got value %d", value, key, actual)
	}
}

func TestEnvironmentConfigSource_GetIntErrNoValueFound(t *testing.T) {
	e := NewEnvironmentConfigSource()

	key := "key"
	env := "TEST_KEY"

	e.Bind(key, env)

	actual, err := e.GetInt(key)
	if err == nil {
		t.Fatalf("expected ErrNoValueFound, got value %d", actual)
	}

	if !errors.Is(err, ErrNoValueFound) {
		t.Fatalf("expected ErrNoValueFound, got error %v", err)
	}
}
