package lrucache

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestClear(t *testing.T) {
	capacity := 2

	expectedCache := createCache(capacity)

	actualCache, err := LRUCache(2)
	if err != nil {
		t.Fatalf("Error in creating cache: %v", err)
	}

	actualCache.Set("0", "a")
	actualCache.Set("1", "b")
	actualCache.Clear()

	assert.Equal(t, expectedCache, actualCache)
}

var testLrucacheTestSuite = []struct {
	name          string
	capacity      int
	expectedCache *Cache
	expectedError error
}{
	{"BasicFuntionality", 1, createCache(1), nil},
	{"No Capacity", 0, createCache(0), errors.New("capacity must be greater than zero")},
}

func TestLRUCache(t *testing.T) {
	for _, testCase := range testLrucacheTestSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actualCache, actualError := LRUCache(testCase.capacity)
			assert.Equal(t, testCase.expectedError, actualError)
			assert.Equal(t, testCase.expectedCache, actualCache)
		})
	}
}
