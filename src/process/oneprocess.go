package process

import "sync"

/*
 * 每个请求
 */
type call struct {
	wg    sync.WaitGroup
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

/*
 * DoProcess处理时间间隔很多的相同的多次请求，合并处理，返回结果，这样能够减轻系统压力
 * 将每次的请求key进行缓存，如果很短时间内有多个相同key的请求到来，则等待，只处理第一次的
 * 其余的直接依据第一次的结果返回
 */
func (g *Group) DoProcess(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()

	if g.mapCall == nil {
		g.mapCall = make(map[string]*call)
	}

	//已经有请求在处理了，则等待
	if c, ok := g.mapCall[key]; ok {
		g.mu.Unlock()
		// fmt.Println("key = " + key)
		c.wg.Wait()
		return c.value, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.mapCall[key] = c

	g.mu.Unlock()

	//处理完请求之后，执行sync.WaitGroup.Done()，这样别的key相同的请求就能直接放回结果
	c.value, c.err = fn()
	c.wg.Done()

	//处理完后删除本次请求的key
	g.mu.Lock()
	delete(g.mapCall, key)
	g.mu.Unlock()

	return c.value, c.err
}
