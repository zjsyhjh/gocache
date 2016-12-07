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
 * ll用来保存所有的item, 采用lru缓存算法, maxEntries指的是最大缓存数量, 0默认是不限制数量
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
	}
}
