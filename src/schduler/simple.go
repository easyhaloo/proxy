package schduler

import "nick.com/proxy/src/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) GetWorkerChan() chan engine.Request {
	return s.WorkerChan
}

func (s *SimpleScheduler) Run() {
	s.WorkerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.WorkerChan <- request
	}()
}
