package lrucache

import (
	"errors"
	"reflect"
	"testing"
)

func TestLRUCache(t *testing.T) {
	capacity := 1

	expectedCache := &Cache{
		Capacity:  1,
		HashMap:   make(map[string]*CacheNode, 1),
		firstNode: nil,
		lastNode:  nil,
	}

	var expectedError error

	actualCache, actualError := LRUCache(capacity)

	if !reflect.DeepEqual(expectedError, actualError) {
		t.Fatalf("\nExpected: %v \nGot: %v", expectedError, actualError)
	}

	if !reflect.DeepEqual(expectedCache, actualCache) {
		t.Fatalf("\nExpected: %v \nGot: %v", expectedCache, actualCache)
	}
}

func TestLRUCache_NoCapacity(t *testing.T) {
	capacity := 0

	var expectedCache *Cache

	expectedError := errors.New("capacity must be greater than zero")

	actualCache, actualError := LRUCache(capacity)

	if !reflect.DeepEqual(expectedError, actualError) {
		t.Fatalf("\nExpected: %v \nGot: %v", expectedError, actualError)
	}

	if !reflect.DeepEqual(expectedCache, actualCache) {
		t.Fatalf("\nExpected: %v \nGot: %v", expectedCache, actualCache)
	}
}

func TestClear(t *testing.T) {
	cache, err := LRUCache(2)
	if err != nil {
		t.Fatalf("Error in creating cache: %v", err)
	}

	cache.Set("0", "a")
	cache.Set("1", "b")
	cache.Clear()

	expectedCache := &Cache{
		Capacity:  2,
		HashMap:   make(map[string]*CacheNode, 2),
		firstNode: nil,
		lastNode:  nil,
	}

	actualCache := cache

	if !reflect.DeepEqual(expectedCache, actualCache) {
		t.Fatalf("\nExpected: %v \nGot: %v", expectedCache, actualCache)
	}
}
