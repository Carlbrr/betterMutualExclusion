package proto

import "sync"

type Cqueue struct {
	queue []int
	lock  sync.RWMutex
}

func (c *Cqueue) Enqueue(name int) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.queue = append(c.queue, name)
}

func (c *Cqueue) Dequeue() {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		c.queue = c.queue[1:]
	}
}

func (c *Cqueue) Front() int {
	if len(c.queue) > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		return c.queue[0]
	}
	return -1
}

func (c *Cqueue) Size() int {
	return len(c.queue)
}

func (c *Cqueue) Empty() bool {
	return len(c.queue) == 0
}
