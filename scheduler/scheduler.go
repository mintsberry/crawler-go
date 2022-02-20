package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) Submit(simple engine.Request) {
	go func() { s.workChan <- simple }()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {
	s.workChan = requests
}
