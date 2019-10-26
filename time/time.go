package time

import (
	"context"
	"sync"
	"time"
)

type clock struct {
	t       time.Time
	mux     sync.RWMutex
	period  time.Duration
	running bool
	cancel  context.CancelFunc
	closed  chan struct{}
}

func newClock(itv time.Duration) *clock {
	ctx, cancel := context.WithCancel(context.Background())
	ck := &clock{t: time.Now(), period: itv, cancel: cancel, closed: make(chan struct{})}
	go ck.run(ctx)
	return ck
}

func (c *clock) run(ctx context.Context) {
	c.mux.Lock()
	if c.running {
		c.mux.Unlock()
		return
	}
	c.t = time.Now()
	c.mux.Unlock()
	ticker := time.NewTicker(c.period)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			close(c.closed)
			return
		case <-ticker.C:
			c.mux.Lock()
			c.t = time.Now()
			c.mux.Unlock()
		}
	}
}

func (c *clock) Stop() {
	c.cancel()
	<-c.closed
}

func (c *clock) GetTime() time.Time {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.t
}

func getTimeFromClock(n int) {
	ck := newClock(time.Millisecond * 100)
	defer ck.Stop()
	for i := 1; i <= n; i++ {
		ck.GetTime()
	}
}

func getTimeNow(n int) {
	for i := 1; i <= n; i++ {
		time.Now()
	}
}
