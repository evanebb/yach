package yach

type ConfigRetriever interface {
	Get(key string) (string, error)
	GetInt(key string) (int, error)
}
