package process

import (
	"sync"
)

/*
 * 每个请求
 */
type call struct {
	wg    *sync.WaitGroup
	value interface{}
	err   error
}

/*
 * 一组类似的请求
 */
type Group struct {
	mu      sync.Mutex
	mapCall map[string]*call
}
