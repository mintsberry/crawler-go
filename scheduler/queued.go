package scheduler

import "crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q QueuedScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q QueuedScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {
	//TODO implement me
	panic("implement me")
}

func (q QueuedScheduler) WorkerReady(requests chan engine.Request) {
	q.workerChan <- requests
}

func (s QueuedScheduler) Run() {
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case s := <-s.workerChan:
				workerQ = append(workerQ, s)
			}
		}
	}()
}
