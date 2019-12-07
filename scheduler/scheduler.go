package scheduler

import "crawler/engine"

type Scheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *Scheduler) WorkerReady(work chan engine.Request) {
	s.workerChan <- work
}

func (s *Scheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *Scheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *Scheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		var reqQueue []engine.Request
		var workerQueue []chan engine.Request

		for{
			var activeReq engine.Request
			var activeWorker chan engine.Request
			if len(reqQueue) > 0 && len(workerQueue) > 0 {
				activeReq = reqQueue[0]
				activeWorker = workerQueue[0]
			}

			select {
				case req := <- s.requestChan:
					reqQueue = append(reqQueue,req)
				case work := <- s.workerChan:
					workerQueue = append(workerQueue,work)
				case activeWorker <- activeReq:
					reqQueue = reqQueue[1:]
					workerQueue = workerQueue[1:]
			}
		}
	}()
}


