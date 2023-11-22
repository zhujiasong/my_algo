package algo

import (
	"my_algo/ds/list"
)

type KVPair struct {
	key, val int
}

type LRUCache struct {
	l   *list.List
	m   map[int]*list.ListNode
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		l:   list.New(nil),
		m:   make(map[int]*list.ListNode, capacity),
		cap: capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.m[key]
	if !ok {
		return -1
	}

	lru.refreshToFront(key, node)
	kv, ok := node.Elem.(*KVPair)
	if !ok {
		panic("kvPair type cast failed")
	}

	return kv.val
}

func (lru *LRUCache) Put(key int, value int) {
	node, ok := lru.m[key]
	// key already exists, just update
	if ok {
		node.Elem = &KVPair{
			key: key,
			val: value,
		}
		lru.refreshToFront(key, node)
		return
	}

	// cache is full, evicts the tail node
	if lru.l.Len() >= lru.cap {
		dropNode := lru.l.Tail()
		lru.l.DeleteNode(dropNode)
		kv, ok := dropNode.Elem.(*KVPair)
		if !ok {
			panic("kvPair type cast failed")
		}
		delete(lru.m, kv.key)
	}

	// insert new node
	node = new(list.ListNode)
	node.Elem = &KVPair{
		key: key,
		val: value,
	}
	lru.l.AddNodeHead(node)
	lru.m[key] = node
}

func (lru *LRUCache) refreshToFront(key int, node *list.ListNode) {
	lru.l.DeleteNode(node)
	lru.l.AddNodeHead(node)
}
