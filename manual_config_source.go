package yach

type ManualConfigSource struct {
	stringBindings map[string]string
	intBindings    map[string]int
}

func NewManualConfigSource() *ManualConfigSource {
	return &ManualConfigSource{
		stringBindings: map[string]string{},
		intBindings:    map[string]int{},
	}
}

// Set will bind the provided string value to the specified key.
func (m *ManualConfigSource) Set(key string, value string) {
	m.stringBindings[key] = value
}

// SetInt will bind the provided integer value to the specified key.
func (m *ManualConfigSource) SetInt(key string, value int) {
	m.intBindings[key] = value
}

func (m *ManualConfigSource) Get(key string) (string, error) {
	value, exists := m.stringBindings[key]
	if exists {
		return value, nil
	}

	return "", ErrNoValueFound
}

func (m *ManualConfigSource) GetInt(key string) (int, error) {
	value, exists := m.intBindings[key]
	if exists {
		return value, nil
	}

	return 0, ErrNoValueFound
}
