package engine

type ConcurrentEngine struct {
	MaxWorkerCount int
	Scheduler      Scheduler
	ItemChan       chan Proxy
	RequestWorker  Processor
}

type Processor func(request Request) (ParseResult, error)

type Scheduler interface {
	Submit(request Request)
	GetWorkerChan() chan Request
	Run()
	Ready
}

type Ready interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seed ...Request) {
	out := make(chan ParseResult, 1024)
	e.Scheduler.Run()
	for i := 0; i < e.MaxWorkerCount; i++ {
		e.createWorker(e.Scheduler.GetWorkerChan(), out, e.Scheduler)
	}

	for _, r := range seed {
		if IsVisited(r.Url) {
			continue
		}
		// 提交任务到调度器
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Proxys {
			go func(proxy Proxy) {
				e.ItemChan <- proxy
			}(item)
		}
		if IsVisited(result.NextRequest.Url) {
			continue
		}
	//	fmt.Println("submit task, the url : " + result.NextRequest.Url)
		e.Scheduler.Submit(result.NextRequest)
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, s Ready) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result, err := e.RequestWorker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
