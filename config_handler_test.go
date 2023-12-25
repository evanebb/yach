package yach

import (
	"testing"
)

func TestConfigHandler_GetSingleConfigSource(t *testing.T) {
	source := NewManualConfigSource()

	key := "key"
	value := "value"
	source.Set(key, value)

	handler := NewConfigHandler()
	handler.Add(source)

	actual, err := handler.Get(key)
	if err != nil {
		t.Fatalf("expected value %s for key %s, got error %v", value, key, err)
	}

	if actual != value {
		t.Fatalf("expected value %s for key %s, got value %s", value, key, actual)
	}
}

func TestConfigHandler_GetIntSingleConfigSource(t *testing.T) {
	source := NewManualConfigSource()

	key := "key"
	value := 1
	source.SetInt(key, value)

	handler := NewConfigHandler()
	handler.Add(source)

	actual, err := handler.GetInt(key)
	if err != nil {
		t.Fatalf("expected value %d for key %s, got error %v", value, key, err)
	}

	if actual != value {
		t.Fatalf("expected value %d for key %s, got value %d", value, key, actual)
	}
}

// TestConfigHandler_GetMultipleConfigSources tests whether using Get() with multiple configuration sources will work.
// This includes values from sources added first being overwritten ones that have been added later.
func TestConfigHandler_GetMultipleConfigSources(t *testing.T) {
	key := "key"
	value := "value"

	manual1 := NewManualConfigSource()
	manual1.Set(key, "thisShouldBeOverwritten")

	manual2 := NewManualConfigSource()
	manual2.Set(key, value)

	handler := NewConfigHandler()
	handler.Add(manual1, manual2)

	actual, err := handler.Get(key)
	if err != nil {
		t.Fatalf("expected value %s for key %s, got error %v", value, key, err)
	}

	if actual != value {
		t.Fatalf("expected value %s for key %s, got value %s", value, key, actual)
	}
}

// TestConfigHandler_GetIntMultipleConfigSources tests whether using GetInt() with multiple configuration sources will work.
// This includes values from sources added first being overwritten ones that have been added later.
func TestConfigHandler_GetIntMultipleConfigSources(t *testing.T) {
	key := "key"
	value := 4321

	manual1 := NewManualConfigSource()
	manual1.SetInt(key, 1234)

	manual2 := NewManualConfigSource()
	manual2.SetInt(key, value)

	handler := NewConfigHandler()
	handler.Add(manual1, manual2)

	actual, err := handler.GetInt(key)
	if err != nil {
		t.Fatalf("expected value %d for key %s, got error %v", value, key, err)
	}

	if actual != value {
		t.Fatalf("expected value %d for key %s, got value %d", value, key, actual)
	}
}

// TestConfigHandler_GetEmptyValue tests for a correct, zero value return value if none of the configuration sources have a value for the requested key.
func TestConfigHandler_GetEmptyValue(t *testing.T) {
	handler := NewConfigHandler()

	actual, err := handler.Get("unknown")
	if err != nil {
		t.Fatalf("expected empty string for unknown key, got error %v", err)
	}

	if actual != "" {
		t.Fatalf("expected empty string for unknown key, got value %s", actual)
	}
}

// TestConfigHandler_GetIntEmptyValue tests for a correct, zero value return value if none of the configuration sources have a value for the requested key.
func TestConfigHandler_GetIntEmptyValue(t *testing.T) {
	handler := NewConfigHandler()

	actual, err := handler.GetInt("unknown")
	if err != nil {
		t.Fatalf("expected empty string for unknown key, got error %v", err)
	}

	if actual != 0 {
		t.Fatalf("expected empty string for unknown key, got value %d", actual)
	}
}
