package utils

import (
	"sync"
)

func SpawnWorkers(n int, worker func()) *sync.WaitGroup {
	wg := &sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			worker()
			wg.Done()
		}()
	}
	return wg
}
