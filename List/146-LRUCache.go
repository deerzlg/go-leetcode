package List

import "container/list"

type entry struct {
	key, value int
}

type LRUCache struct {
	keyToNode map[int]*list.Element
	list      *list.List
	capacity  int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		keyToNode: make(map[int]*list.Element),
		list:      list.New(),
		capacity:  capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.keyToNode[key]; ok {
		lru.list.MoveToFront(node)
		return node.Value.(entry).value
	} else {
		return -1
	}
}

func (lru *LRUCache) Put(key, value int) {
	newEntry := entry{key, value}
	if node, ok := lru.keyToNode[key]; ok {
		node.Value = newEntry
		lru.list.MoveToFront(node)
	} else {
		newNode := lru.list.PushFront(newEntry)
		lru.keyToNode[key] = newNode
		if lru.list.Len() > lru.capacity {
			delete(lru.keyToNode, lru.list.Remove(lru.list.Back()).(entry).key)
		}
	}
}
