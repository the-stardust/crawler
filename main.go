package main

import (
	"crawler/engine"
	"crawler/parser"
	"crawler/persist"
	"crawler/scheduler"
	"github.com/olivere/elastic/v7"
)

func main() {
	client,err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		)
	if err != nil {
		panic(err)
	}
	itemChan,err := persist.ItemSaver(client,"zhenai_profile")
	if err != nil {
		panic(err)
	}
	e := engine.QueueEngine{
		Scheduler:   &scheduler.Scheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})
}
