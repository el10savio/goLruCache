package lrucache

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	{"LargeCache", 10000000, createCache(10000000), nil},
	{"No Capacity", 0, createCache(0), errors.New("capacity must be greater than zero")},
}

func TestLRUCache(t *testing.T) {
	for _, testCase := range testLrucacheTestSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actualCache, actualError := LRUCache(testCase.capacity)
			defer clearCaches([]*Cache{testCase.expectedCache, actualCache})

			assert.Equal(t, testCase.expectedError, actualError)
			assert.Equal(t, testCase.expectedCache, actualCache)
		})
	}
}

var testGetTestSuite = []struct {
	name          string
	baseCache     *Cache
	key           string
	expectedValue interface{}
}{
	{"BasicFuntionality", fillCache(3, []insertTuple{{"a", 1}, {"b", 2}, {"c", 3}}), "b", 2},
	{"DuplicateElements", fillCache(1, []insertTuple{{"a", 1}, {"a", 1}, {"a", 1}}), "a", 1},
	{"RewrittenElements", fillCache(1, []insertTuple{{"a", 1}, {"a", 2}, {"a", 3}}), "a", 3},
	{"ElementNotPresent", fillCache(3, []insertTuple{{"a", 1}, {"b", 2}, {"c", 3}}), "z", nil},
	{"SingleElement", fillCache(1, []insertTuple{{"a", 1}}), "a", 1},
	{"EmptyCache", fillCache(3, []insertTuple{}), "b", nil},
}

func TestGet(t *testing.T) {
	for _, testCase := range testGetTestSuite {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			// t.Parallel()

			cache := testCase.baseCache
			defer clearCaches([]*Cache{cache})

			actualValue := cache.Get(testCase.key)

			assert.Equal(t, testCase.expectedValue, actualValue)
		})
	}
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

type insertTuple struct {
	Key   string
	Value interface{}
}

func fillCache(capacity int, elements []insertTuple) *Cache {
	cache := createCache(capacity)
	if cache == nil {
		return nil
	}

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
