package scheduler

import "go-demo-reptile/bussiness/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request

	workerChan chan chan engine.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) Submit(r engine.Request) {

	q.requestChan <- r
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) MastWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (q *QueuedScheduler) Run() {

	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)

	go func() {
		var rQ []engine.Request
		var wQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(rQ) > 0 && len(wQ) > 0 {
				activeRequest = rQ[0]
				activeWorker = wQ[0]
			}

			select {
			case r := <-q.requestChan:
				rQ = append(rQ, r)
			case w := <-q.workerChan:
				wQ = append(wQ, w)

			case activeWorker <- activeRequest:

				rQ = rQ[1:]
				wQ = wQ[1:]

			}

		}

	}()
}
