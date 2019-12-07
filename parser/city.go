package parser

import (
	"crawler/engine"
	"log"
	"regexp"
)

var cityRegex = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^<]+)</a>`)

func ParserCity(content []byte,_ string)(result engine.ParserResult){

	matches := cityRegex.FindAllSubmatch(content,-1)

	for _,val := range matches{
		log.Printf("got user : %s",string(val[2]))
		result.Requests = append(
			result.Requests,
			engine.Request{
				Url:        string(val[1]),
				ParserFunc: ProfileParser(string(val[2])),
			})
	}

	return
}

func ProfileParser(name string)engine.ParserFunc{
	return func(content []byte, url string) engine.ParserResult {
		return ParserProfile(content,url,name)
	}
}