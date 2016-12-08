package process

import "sync"

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

func (g *Group) DoProcess(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()

	if g.mapCall == nil {
		g.mapCall = make(map[string]*call)
	}

	//已经有请求在处理了，则等待
	if c, ok := g.mapCall[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.value, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.mapCall[key] = c

	g.mu.Unlock()

	c.value, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.mapCall, key)
	g.mu.Unlock()

	return c.value, c.err
}
