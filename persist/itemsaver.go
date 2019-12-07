package persist

import (
	"context"
	"crawler/engine"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver(client *elastic.Client,index string)(chan engine.Item,error){
	out := make(chan engine.Item)
	count := 0
	go func() {
		for {
			item := <- out
			err := save(client,index,item)
			if err != nil {
				log.Printf("save item error : %v",err)
				continue
			}
			count++
			log.Printf("save item success, count : %d  item : %v",count,item)
		}
	}()
	return out,nil
}

func save(client *elastic.Client,index string,item engine.Item)error{
	_,err := client.Index().Index(index).Id(item.Id).BodyJson(item).Do(context.Background())
	return err
}
