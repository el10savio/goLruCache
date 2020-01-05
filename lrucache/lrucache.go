package lrucache

import (
	"errors"
)

// Cache specifies the capacity of the cache
// It consists of a HashMap with an integer key
// corresponding to a particular cache node
// Also the first & last nodes are available
type Cache struct {
	Capacity int
	HashMap  map[string]*CacheNode

	firstNode *CacheNode
	lastNode  *CacheNode

	CacheFunctions
}

// CacheNode consists of a key value pair.
// It is implemented in a doubly linked list
// format for easy removal & update so
// we also store the previous & next
// cache nodes
type CacheNode struct {
	Key   string
	Value interface{}

	previousNode *CacheNode
	nextNode     *CacheNode
}

// CacheFunctions associated functions are:
// Get(key) to obtain a particular node
// Set(key, value) to place it into the cache
type CacheFunctions interface {
	Get(key string) interface{}
	Set(key string, value interface{})
}

// Get first checks if the key is present
// in the map and then returns it.
// In doing so we have used this node
// and thus push it to the first node
func (cache *Cache) Get(key string) interface{} {
	node := cache.HashMap[key]
	if node != nil {
		cache.removeNode(key)
		cache.pushToHead(key, node.Value)
		return node.Value
	}
	return nil
}

// Set first removes any node with the
// same key and then adds the given
// key value data as a new node
// to the head
func (cache *Cache) Set(key string, value interface{}) {
	cache.removeNode(key)
	cache.pushToHead(key, value)

	// Check if capacity exceeds
	if len(cache.HashMap) > cache.Capacity {
		cache.removeNode(cache.lastNode.Key)
	}

}

// pushToHead() adds the given node to the
// cache. Then, it checks if the capacity
// is more than specificed and if so
// removes the last cache node
func (cache *Cache) pushToHead(key string, value interface{}) {
	// Connect the given node
	// to the first node
	node := &CacheNode{
		Key:          key,
		Value:        value,
		nextNode:     cache.firstNode,
		previousNode: nil,
	}

	// Push the previous first node
	// to after the given node
	if cache.firstNode != nil {
		cache.firstNode.previousNode = node
	}

	// If only the given node is
	// present, it is the
	// lastNode also
	if cache.lastNode == nil {
		cache.lastNode = node
	}

	// Make the given node the first node
	cache.firstNode = node

	// Add the node to the HashMap
	cache.HashMap[key] = node
}

// removeNode() removes the given node from the
// cache by unlinking it from its previous
// and next nodes and then linking each
// other
func (cache *Cache) removeNode(key string) {
	// Check if node exists
	node := cache.HashMap[key]
	if node != nil {
		// If there is a node behind
		// the given node
		if node.previousNode != nil {
			(node.previousNode).nextNode = node.nextNode
		}

		// If there is a node after
		// the given node
		if node.nextNode != nil {
			(node.nextNode).previousNode = node.previousNode
		} else {
			cache.lastNode = node.previousNode
		}

		// Remove the node from the HashMap
		delete(cache.HashMap, key)
	}
	return
}

// Clear removes all nodes from the cache
func (cache *Cache) Clear() {
	cache.HashMap = make(map[string]*CacheNode, cache.Capacity)
	cache.firstNode = nil
	cache.lastNode = nil
}

// LRUCache creates a new empty LRUCache
func LRUCache(capacity int) (*Cache, error) {
	// Validate the capacity
	if capacity > 0 {
		cache := &Cache{
			Capacity:  capacity,
			HashMap:   make(map[string]*CacheNode, capacity),
			firstNode: nil,
			lastNode:  nil,
		}
		return cache, nil
	}
	return nil, errors.New("capacity must be greater than zero")
}
