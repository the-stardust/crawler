package engine

import "log"

type QueueEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan Item
}

func (e *QueueEngine) Run (seeds ...Request){
	e.Scheduler.Run()
	out := make(chan ParserResult)
	for i := 0; i < e.WorkerCount; i++{
		createWorker(e.Scheduler.WorkChan(),out,e.Scheduler)
	}

	for _,seed := range seeds{
		e.Scheduler.Submit(seed)
	}

	for{
		result := <- out

		for _,val := range result.Items{
			go func() {
				e.ItemChan <- val
			}()
		}

		for _, req := range result.Requests{
			e.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request,out chan ParserResult,notify ReadyNotify){

	go func() {
		for{
			notify.WorkerReady(in)
			req := <- in
			result,err := worker(req)
			if err != nil {
				log.Printf("worker err : %v",err)
				continue
			}
			out <- result
		}
	}()
}



