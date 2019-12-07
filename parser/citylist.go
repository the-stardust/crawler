package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParserCityList(content []byte,_ string)(result engine.ParserResult){

	matches := cityListRe.FindAllSubmatch(content,-1)
	//limit := 10
	for _,val := range matches{
		log.Printf("got city : %s",string(val[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(val[1]),
				ParserFunc: ParserCity,
			})
		//limit--
		//if limit == 0{
		//	break
		//}
	}

	return
}
