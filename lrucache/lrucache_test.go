package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClear(t *testing.T) {
	for _, testCase := range testClearTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			cache := testCase.baseCache
			cache.Clear()

			assert.Equal(t, testCase.expectedCache, cache)
		})
	}

}

func TestLRUCache(t *testing.T) {
	for _, testCase := range testLrucacheTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			actualCache, actualError := LRUCache(testCase.capacity)
			defer clearCaches([]*Cache{testCase.expectedCache, actualCache})

			assert.Equal(t, testCase.expectedError, actualError)
			assert.Equal(t, testCase.expectedCache, actualCache)
		})
	}
}

func TestSet(t *testing.T) {
	for _, testCase := range testSetTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			cache := testCase.baseCache
			defer clearCaches([]*Cache{cache})

			actualHead := cache.getHead()

			assert.Equal(t, testCase.expectedTuple.Key, actualHead.Key)
			assert.Equal(t, testCase.expectedTuple.Value, actualHead.Value)
		})
	}
}

func TestGet(t *testing.T) {
	for _, testCase := range testGetTestSuite {
		t.Run(testCase.name, func(t *testing.T) {
			cache := testCase.baseCache
			defer clearCaches([]*Cache{cache})

			actualValue := cache.Get(testCase.key)

			assert.Equal(t, testCase.expectedValue, actualValue)
		})
	}
}
