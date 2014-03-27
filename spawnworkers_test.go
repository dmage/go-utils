package utils

import "testing"

func TestSpawnWorkers(t *testing.T) {
	inChan := make(chan int)
	outChan := make(chan int)

	sum := 0
	collector := SpawnWorkers(1, func() {
		for i := range outChan {
			sum += i
		}
	})

	SpawnWorkers(3, func() {
		for i := range inChan {
			outChan <- i * i
		}
	}).Defer(func() {
		close(outChan)
	})

	for i := 0; i < 5; i++ {
		inChan <- i
	}
	close(inChan)

	collector.Wait()

	want := 0*0 + 1*1 + 2*2 + 3*3 + 4*4
	if sum != want {
		t.Errorf("sum = %v, want = %v\n", sum, want)
	}
}
