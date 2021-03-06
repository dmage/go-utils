package utils

import (
	"sync"
)

type Workers struct {
	wg sync.WaitGroup
}

// Spawn n goroutines and return Workers object
func SpawnWorkers(n int, worker func()) *Workers {
	workers := &Workers{}
	for i := 0; i < n; i++ {
		workers.wg.Add(1)
		go func() {
			worker()
			workers.wg.Done()
		}()
	}
	return workers
}

// Execute function after all workers done their work
func (workers *Workers) Defer(f func()) *Workers {
	go func() {
		workers.wg.Wait()
		f()
	}()
	return workers
}

// Wait for the completion of all workers
func (workers *Workers) Wait() {
	workers.wg.Wait()
}
