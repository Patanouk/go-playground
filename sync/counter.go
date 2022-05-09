package sync

import "sync"

type Counter struct {
	value int
	lock  sync.Mutex
}

func (c *Counter) Inc() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.value += 1
}

func (c *Counter) Value() int {
	return c.value
}
