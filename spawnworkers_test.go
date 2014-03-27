package utils_test

import (
	"fmt"
	. "github.com/dmage/go-utils"
)

func ExampleSpawnWorkers() {
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

	fmt.Println(sum) // 0*0 + 1*1 + 2*2 + 3*3 + 4*4 == 30

	// Output: 30
}
