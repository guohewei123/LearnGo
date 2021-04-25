package engine

import (
	"crawler/model"
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
	ReadyNotifier
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParserResult)
	e.Scheduler.Run()

	for i :=0; i <e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	profileCount := 0
	for true {
		result := <- out
		for _, item := range result.Items {
			if _, ok := item.(model.Profile); ok {
				log.Printf("Got profile #%d: %v", profileCount, item)
				profileCount++
			}
		}
		// URL dedup 去重
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				//log.Printf("duplicate request: %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for true {
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedUrls = make(map[string]bool)
func isDuplicate(r string) bool {
	if visitedUrls[r] {
		return true
	}
	visitedUrls[r] = true
	return false
}