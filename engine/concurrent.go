package engine

import "fmt"

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)
	for i := 0; i < e.WorkCount; i++ {
		createWorker(in, out)
	}

	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}
	count := 0
	for {
		result := <-out //wait
		for _, item := range result.Items {
			count++
			fmt.Printf("Got item：%s, count: %d\n", item, count)
		}
		for _, request := range result.Request {
			request := request
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
