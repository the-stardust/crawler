package engine

type Request struct {
	Url string
	ParserFunc ParserFunc
}
type Item struct {
	Id string
	Url string
	Payload interface{}
}

type ParserResult struct {
	Requests []Request
	Items []Item
}

type ParserFunc func(content []byte,url string)ParserResult

type Scheduler interface {
	ReadyNotify
	Submit(request Request)
	Run()
	WorkChan()chan Request
}

type ReadyNotify interface {
	WorkerReady(chan Request)
}