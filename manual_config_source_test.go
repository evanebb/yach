package yach

import (
	"errors"
	"testing"
)

func TestManualConfigSource_Get(t *testing.T) {
	m := NewManualConfigSource()

	key := "key"
	value := "value"
	m.Set(key, value)

	actual, err := m.Get(key)
	if err != nil {
		t.Fatalf("expected value %s for key %s, got error %v", value, key, err)
	}

	if actual != value {
		t.Fatalf("expected value %s for key %s, got value %s", value, key, actual)
	}
}

func TestManualConfigSource_GetErrNoValueFound(t *testing.T) {
	m := NewManualConfigSource()

	actual, err := m.Get("unknown")
	if err == nil {
		t.Fatalf("expected ErrNoValueFound for unknown key, got value %s", actual)
	}

	if !errors.Is(err, ErrNoValueFound) {
		t.Fatalf("expected ErrNoValueFound for unknown key, got error %v", err)
	}
}

func TestManualConfigSource_GetInt(t *testing.T) {
	m := NewManualConfigSource()

	key := "key"
	value := 1
	m.SetInt(key, value)

	actual, err := m.GetInt(key)
	if err != nil {
		t.Fatalf("expected value %d for key %s, got error %v", value, key, err)
	}

	if actual != value {
		t.Fatalf("expected value %d for key %s, got value %d", value, key, actual)
	}
}

func TestManualConfigSource_GetIntErrNoValueFound(t *testing.T) {
	m := NewManualConfigSource()

	actual, err := m.GetInt("unknown")
	if err == nil {
		t.Fatalf("expected ErrNoValueFound for unknown key, got value %d", actual)
	}

	if !errors.Is(err, ErrNoValueFound) {
		t.Fatalf("expected ErrNoValueFound for unknown key, got error %v", err)
	}
}
