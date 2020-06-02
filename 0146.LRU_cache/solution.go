package leetcode0146

// DoublyLinkedList : 双向链表
type DoublyLinkedList struct {
	prev, next *DoublyLinkedList
	key, value int
}

// LRUCache : LRU cache
type LRUCache struct {
	cap, size  int
	cache      map[int]*DoublyLinkedList
	head, tail *DoublyLinkedList
}

// Constructor : init a LRUCache with capacity
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap: capacity,
		head: &DoublyLinkedList{key: 0, value: 0},
		tail: &DoublyLinkedList{key: 0, value: 0},
		cache: map[int]*DoublyLinkedList{},
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

// Get : LRUCache get a node
func (c *LRUCache) Get(key int) int {
	if _, ok := c.cache[key]; !ok {
		return -1
	}
	node := c.cache[key]
	// remove node
	node.prev.next = node.next
	node.next.prev = node.prev

	// add to head
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
	return node.value
}

// Put : LRUCache put a node
func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.cache[key]; !ok {
		node := &DoublyLinkedList{key: key, value: value}
		c.cache[key] = node
		// add to head
		node.prev = c.head
		node.next = c.head.next
		c.head.next.prev = node
		c.head.next = node
		c.size++
		if c.size > c.cap {
			// remove the tail node
			node := c.tail.prev
			node.prev.next = node.next
			node.next.prev = node.prev
			delete(c.cache, node.key)
			c.size--
		}
	} else {
		node := c.cache[key]
		// change value
		node.value = value
		// remove node
		node.prev.next = node.next
		node.next.prev = node.prev
		// add to head
		node.prev = c.head
		node.next = c.head.next
		c.head.next.prev = node
		c.head.next = node
	}
}
