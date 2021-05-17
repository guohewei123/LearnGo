package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParserResult, error)

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

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for true {
		result := <-out
		for _, item := range result.Items {
			//log.Printf("------Item Saver: got item %v", item)
			temp := item
			go func() {
				e.ItemChan <- temp
			}()
			//if _, ok := item.(model.Profile); ok {
			//
			//}
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

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for true {
			ready.WorkerReady(in)
			request := <-in
			//result, err := Worker(request)
			result, err := e.RequestProcessor(request)
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
