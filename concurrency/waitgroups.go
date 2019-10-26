package concurrency

import (
	"sync"
	"time"
)

func spawnGoSleep(n int, dur time.Duration) {
	wg := &sync.WaitGroup{}
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go func() {
			time.Sleep(dur)
			wg.Done()
		}()
	}
	wg.Wait()
}

func spawnGoSleepWG(n int, dur time.Duration) {
	wg := &sync.WaitGroup{}
	wg.Add(n)
	for i := 1; i <= n; i++ {
		go sleepDoneWG(wg, dur)
	}
	wg.Wait()
}

func sleepDoneWG(wg *sync.WaitGroup, dur time.Duration) {
	time.Sleep(dur)
	wg.Done()
}
