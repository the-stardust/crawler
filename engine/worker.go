package engine

import (
	"crawler/fetcher"
	"log"
)

func worker(r Request)(result ParserResult,err error){

	content,err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch url error %s : %v",r.Url,err)
		return
	}
	return r.ParserFunc(content,r.Url),nil
}
