package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}
type KeyValue struct {
	key   int
	value int
}

func newLRUCache(cap int) *LRUCache {
	return &LRUCache{
		// Initialkize the cache with the given capacity
		// and create a new list to store the key-value pairs
		// The list will be used to keep track of the order
		// of the key-value pairs, with the most recently used
		// pairs at the front and the least recently used pairs
		// at the back
		capacity: cap,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

//pointer for LRUCache becaus we want to modify the original struct
// and not a copy of it
// In the put method we are just reading so pass by value is fine
func (lru *LRUCache) Put(kv KeyValue) {
	// Check if the key already exists
	// If it exists, move the element to the front of the list
	// and update the value
	// If it doesn't exist, add it to the front of the list
	// and add it to the cache
	if elem, exists := lru.cache[kv.key]; exists {
		lru.list.MoveToFront(elem)
		elem.Value = kv
		fmt.Println("Updated:", kv.key, "to", kv.value)
		return
	}
	// If the cache is full, evict the least recently used item
	if len(lru.cache) == lru.capacity {
		lru.Evict()
	}
	// Add the new key-value pair to the front of the list
	elem := lru.list.PushFront(kv)
	fmt.Println("Added:", kv.key, "with value", kv.value)
	// Store the element in the cache
	// using the key as the map key
	lru.cache[kv.key] = elem

}

func (lru *LRUCache) Get(key int) int {
	// Check if the key exists in the cache
	// If it exists, move the element to the front of the list
	// and return the value
	// If it doesn't exist, return -1
	if elem, ok := lru.cache[key]; ok {
		// Move the accessed element to the front of the list
		// to mark it as recently used
		// and return the value
		lru.list.MoveToFront(elem)
		return elem.Value.(KeyValue).value
	}
	return -1
}

func (lru *LRUCache) Evict() {
	// Evict the least recently used item
	// which is at the back of the list
	// Remove it from the list and the cache
	// Get the last element in the list
	// and remove it from the cache
	// and the list
	// The last element in the list is the least recently used
	// item, so we remove it from the list and the cache
	// and return
	// the value of the least recently used item
	// to the caller
	lastElem := lru.list.Back()
	if lastElem != nil {
		kv := lastElem.Value.(KeyValue).key
		delete(lru.cache, kv)
		lru.list.Remove(lastElem)
	}
}

func (lru *LRUCache) PrintCache() {
	// Print the current state of the cache
	// Iterate through the list and print the key-value pairs
	// The most recently used pairs will be at the front
	// and the least recently used pairs will be at the back

	for e := lru.list.Front(); e != nil; e = e.Next() {
		kv := e.Value.(KeyValue)
		fmt.Printf("(%d:%d) ", kv.key, kv.value)
	}
	fmt.Println()
}

func main() {
	lruCache := newLRUCache(2)
	lruCache.Put(KeyValue{key: 1, value: 1})
	lruCache.Put(KeyValue{key: 2, value: 2})
	fmt.Println("Get 1:", lruCache.Get(1))   // Should print 1
	lruCache.Put(KeyValue{key: 3, value: 3}) // Evicts key 2
	fmt.Println("Get 2:", lruCache.Get(2))   // Should print -1
	lruCache.PrintCache()
	lruCache.Put(KeyValue{key: 4, value: 4}) // Evicts key 1
	fmt.Println("Get 1:", lruCache.Get(1))   // Should print -1
	lruCache.PrintCache()
}
