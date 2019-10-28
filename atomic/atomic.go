package atomic

import (
	"sync"
	"sync/atomic"
)

type atomicCounter struct {
	n uint64
}

type muxCounter struct {
	n   uint64
	mux sync.Mutex
}

func (c *atomicCounter) increment(n uint64) {
	atomic.AddUint64(&c.n, n)
}

func (c *muxCounter) increment(n uint64) {
	c.mux.Lock()
	c.n += n
	c.mux.Unlock()
}
