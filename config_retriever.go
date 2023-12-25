package yach

// ConfigRetriever is the interface that needs to be implemented by every type that should be used as a configuration source for ConfigHandler.
type ConfigRetriever interface {
	// Get will look up the value for the passed key from the configuration source, and return it as a string.
	Get(key string) (string, error)
	// GetInt will look up the value for the passed key from the configuration source, and return it as an integer.
	GetInt(key string) (int, error)
}
