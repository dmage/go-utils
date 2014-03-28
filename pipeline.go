package utils

type PipelineFirst struct {
	N      int
	Worker func(outChan chan<- interface{})
}

type PipelineMiddle struct {
	N      int
	Worker func(inChan <-chan interface{}, outChan chan<- interface{})
}

type PipelineLast struct {
	Worker func(inChan <-chan interface{})
}

type Pipeline struct {
	First  PipelineFirst
	Middle []PipelineMiddle
	Last   PipelineLast
}

func (p *Pipeline) Run() *Workers {
	inChan := make(chan interface{})
	var outChan chan interface{}

	func(inChan <-chan interface{}) {
		SpawnWorkers(1, func() {
			p.Last.Worker(inChan)
		})
	}(inChan)

	for i := len(p.Middle) - 1; i >= 0; i-- {
		outChan = inChan
		inChan = make(chan interface{})
		func(i int, inChan <-chan interface{}, outChan chan<- interface{}) {
			SpawnWorkers(p.Middle[i].N, func() {
				p.Middle[i].Worker(inChan, outChan)
			}).Defer(func() {
				close(outChan)
			})
		}(i, inChan, outChan)
	}

	outChan = inChan
	last := SpawnWorkers(p.First.N, func() {
		p.First.Worker(outChan)
	}).Defer(func() {
		close(outChan)
	})
	return last
}
