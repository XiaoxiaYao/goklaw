package engine

import (
	"fmt"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	GetWorkerChan() chan Request
	Run()
	ReadyNotifier
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (ce *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	ce.Scheduler.Run()

	for i := 0; i < ce.WorkerCount; i++ {
		createWorker(ce.Scheduler.GetWorkerChan(), out, ce.Scheduler)
	}

	for _, r := range seeds {
		ce.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: #%d: %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			ce.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			fmt.Println(len(in))
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
