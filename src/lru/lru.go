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
 * ll用来保存所有的item, 采用lru缓存算法, maxEntries指的是最大缓存数量
 * list.Element指向的类型为entry类型
 */
type Cache struct {
	ll         *list.List
	cache      map[interface{}]*list.Element
	maxEntries int
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
