package lru

import "container/list"

type Item struct {
	key, value int
}

type LRUCache struct {
	cap   int
	ll    *list.List
	items map[int]*list.Element
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity, ll: list.New(), items: make(map[int]*list.Element, capacity)}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.items[key]; ok {
		c.ll.MoveToFront(node)
		item := node.Value.(Item)
		return item.value
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.items[key]; ok {
		node.Value = Item{key: key, value: value}
		c.ll.MoveToFront(node)
		return
	}

	if len(c.items) == c.cap {
		lastNode := c.ll.Back()
		itemToRemove := lastNode.Value.(Item)
		c.ll.Remove(lastNode)
		delete(c.items, itemToRemove.key)
	}

	node := c.ll.PushFront(Item{key: key, value: value})
	c.items[key] = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
