package lrucache

type tuple struct {
	Key   string
	Value interface{}
}

func createCache(capacity int) *Cache {
	if capacity == 0 {
		return nil
	}

	return &Cache{
		Capacity:  capacity,
		HashMap:   make(map[string]*CacheNode, capacity),
		firstNode: nil,
		lastNode:  nil,
	}
}

func fillCache(capacity int, elements []tuple) *Cache {
	cache := createCache(capacity)

	for _, element := range elements {
		cache.Set(element.Key, element.Value)
	}

	return cache
}

func clearCaches(caches []*Cache) {
	for _, cache := range caches {
		if cache != nil {
			cache.Clear()
		}
	}
}
