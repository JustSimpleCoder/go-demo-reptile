package engine

import ()

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	WorkCount int
	Scheduler Scheduler
	ItemChan  chan interface{}
}

func (ce *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	ce.Scheduler.Run()

	for i := 0; i < ce.WorkCount; i++ {

		createWork(ce.Scheduler.WorkerChan(), out, ce.Scheduler)
	}

	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}
	for {
		result := <-out
		for _, item := range result.Items {

			go func() { ce.ItemChan <- item }()
		}

		for _, r := range result.Requests {
			ce.Scheduler.Submit(r)
		}
	}

}

func createWork(in chan Request, out chan ParseResult, ready ReadyNotifier) {

	go func() {
		for {
			ready.WorkerReady(in)
			r := <-in
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
