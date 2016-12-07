package lru

import (
	"container/list"
)

type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

/*
 * map表用于保存所有的item，ll采用lru缓存算法对item进行存储，每次剔除都剔除链表的最后一个item
 * maxEntries指的是最大缓存数量, 0默认是不限制数量
 * list.Element指向的类型为entry类型
 */
type Cache struct {
	ll         *list.List
	cache      map[interface{}]*list.Element
	maxEntries int
	Evict      func(key Key, value interface{})
}

/*
 * create a Cache
 */
func NewCache(maxEntries int) *Cache {
	return &Cache{
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
		maxEntries: maxEntries,
	}
}

/*
 * add value to Cache
 */
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil {
		c.ll = list.New()
		c.cache = make(map[interface{}]*list.Element)
	}

	if e, ok := c.cache[key]; ok {
		c.ll.MoveToFront(e)
		e.Value.(*entry).value = value
		return
	}
	e := c.ll.PushFront(&entry{key, value})
	c.cache[key] = e
	if c.maxEntries != 0 && c.ll.Len() > c.maxEntries {
		c.RemoveOldest()
	}
}

/*
 * Get value from Cache
 */
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}
	if e, hit := c.cache[key]; hit {
		c.ll.MoveToFront(e)
		return e.Value.(*entry).value, true
	}
	return
}

/*
 * 根据LRU算法移除链表最末尾的item
 */
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	e := c.ll.Back()
	if e != nil {
		c.removeElement(e)
	}
}

/*
 * remove item from list and Cache
 */
func (c *Cache) removeElement(e *list.Element) {
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)
	if c.Evict != nil {
		c.Evict(kv.key, kv.value)
	}
}
