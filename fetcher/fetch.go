package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var rateLimit = time.Tick(time.Millisecond * 100)
func Fetch(url string) (content []byte,err error){
	<- rateLimit
	client := &http.Client{}

	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		log.Printf("fetch url error url : %s error : %v",url,err)
		return
	}
	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	req.Header.Set("Referer","http://www.zhenai.com/zhenghun/zhengzhou")

	resp,err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("client do error : ", err)
		return nil,err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("status  code error : %v", resp.StatusCode)
		return nil,err
	}
	temp := resp.Body
	e := determineEncoding(temp)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader)encoding.Encoding{
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e
}
