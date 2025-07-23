package lru

import "container/list"

type Cache[K comparable, V any] struct {
	cap   int
	items map[K]*list.Element
	ll    *list.List
}

type Item[K comparable, V any] struct {
	key   K
	value V
}

func New[K comparable, V any](capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		cap:   capacity,
		items: make(map[K]*list.Element, capacity),
		ll:    list.New(),
	}
}

func (c *Cache[K, V]) Get(key K) (res V, ok bool) {
	if node, exists := c.items[key]; exists {
		c.ll.MoveToFront(node)
		res = node.Value.(Item[K, V]).value
		return res, true
	}

	return
}

func (c *Cache[K, V]) Add(key K, value V) {
	if node, ok := c.items[key]; ok {
		newItem := Item[K, V]{key: key, value: value}
		node.Value = newItem
		c.ll.MoveToFront(node)

		return
	}

	if len(c.items) == c.cap {
		nodeToRemove := c.ll.Back()
		item := nodeToRemove.Value.(Item[K, V])
		delete(c.items, item.key)
		c.ll.Remove(nodeToRemove)
	}

	node := c.ll.PushFront(Item[K, V]{key: key, value: value})
	c.items[key] = node
}
