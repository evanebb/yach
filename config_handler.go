package yach

type ConfigHandler struct {
	retrievers []ConfigRetriever
}

func NewConfigHandler() ConfigHandler {
	return ConfigHandler{}
}

// Add adds the specified sources implementing the ConfigRetriever interface onto the retriever stack
func (c *ConfigHandler) Add(retriever ...ConfigRetriever) {
	c.retrievers = append(c.retrievers, retriever...)
}

func (c *ConfigHandler) Get(key string) (string, error) {
	var value string

	for _, r := range c.retrievers {
		tempValue, err := r.Get(key)
		if err != nil {
			continue
		}

		value = tempValue
	}

	return value, nil
}

func (c *ConfigHandler) GetInt(key string) (int, error) {
	var value int

	for _, r := range c.retrievers {
		tempValue, err := r.GetInt(key)
		if err != nil {
			continue
		}

		value = tempValue
	}

	return value, nil
}
