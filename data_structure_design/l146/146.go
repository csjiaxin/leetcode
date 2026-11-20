package l146

import "container/list"

type entry struct{ key, value int }

// LRUCache uses the list front for the most recently accessed key.
// Construct before use; do not copy after use. Not safe for concurrent use.
// Get and Put take average O(1) time; space is O(capacity).
type LRUCache struct {
	capacity int
	order    *list.List
	index    map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	if capacity < 0 {
		panic("negative capacity")
	}
	return LRUCache{capacity, list.New(), make(map[int]*list.Element)}
}

func (c *LRUCache) Get(key int) int {
	element, ok := c.index[key]
	if !ok {
		return -1
	}
	c.order.MoveToFront(element)
	return element.Value.(entry).value
}

func (c *LRUCache) Put(key int, value int) {
	if element, ok := c.index[key]; ok {
		element.Value = entry{key, value}
		c.order.MoveToFront(element)
		return
	}
	if c.capacity == 0 {
		return
	}
	c.index[key] = c.order.PushFront(entry{key, value})
	if c.order.Len() > c.capacity {
		oldest := c.order.Back()
		delete(c.index, oldest.Value.(entry).key)
		c.order.Remove(oldest)
	}
}
