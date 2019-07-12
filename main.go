package main

import (
	"fmt"

	LRU "./lrucache"
)

func main() {

	// Create LRU Cache with capacity of 5
	cache, err := LRU.LRUCache(5)
	if err != nil {
		panic(err)
	}

	fmt.Println("Cache:", cache)
	fmt.Println("")

	// Push data to cache
	cache.Set("0", "a")
	fmt.Println("Cache:", cache)
	cache.Set("1", "b")
	fmt.Println("Cache:", cache)
	cache.Set("2", "c")
	fmt.Println("Cache:", cache)
	cache.Set("3", "d")
	fmt.Println("Cache:", cache)
	cache.Set("4", "e")
	fmt.Println("Cache:", cache)
	fmt.Println("")

	// Get data from the cache
	cache.Get("0")
	fmt.Println("Cache:", cache)
	cache.Get("1")
	fmt.Println("Cache:", cache)
	fmt.Println("")

	// Set new values to old data
	cache.Set("2", "x")
	fmt.Println("Cache:", cache)
	cache.Set("1", "f")
	fmt.Println("Cache:", cache)
	fmt.Println("")

	// Push new data to discard old data
	cache.Set("5", "q")
	fmt.Println("Cache:", cache)
	cache.Set("6", "w")
	fmt.Println("Cache:", cache)
	cache.Set("7", "e")
	fmt.Println("Cache:", cache)
	cache.Set("8", "r")
	fmt.Println("Cache:", cache)
	cache.Set("9", "t")
	fmt.Println("Cache:", cache)
	fmt.Println("")

	// Clear the cache
	cache.Clear()
	fmt.Println("Cache:", cache)
	fmt.Println("")

}
