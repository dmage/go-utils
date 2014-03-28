package utils_test

import (
	"fmt"
	. "github.com/dmage/go-utils"
)

func ExamplePipeline() {
	p := Pipeline{
		PipelineFirst{1, func(out chan<- interface{}) {
			for i := 0; i < 5; i++ {
				out <- i
			}
		}},
		[]PipelineMiddle{
			{3, func(in <-chan interface{}, out chan<- interface{}) {
				for i := range in {
					out <- i.(int) * i.(int)
				}
			}},
		},
		PipelineLast{func(in <-chan interface{}) {
			sum := 0
			for i := range in {
				sum = sum + i.(int)
			}
			fmt.Println(sum) // 0*0 + 1*1 + 2*2 + 3*3 + 4*4 == 30
		}},
	}
	p.Run().Wait()

	// Output: 30
}
