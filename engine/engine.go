package engine

import (
	"github.com/scott-x/crawler/fetcher"
	"log"
)

func Run(seeds ...Request){
	var requests []Request
	for _, r := range seeds{
		requests = append(requests,r)
	}

	for len(requests)>0{
		r := requests[0]
		requests = requests[1:]
		body,err := fetcher.Fetch(r.Url)
		if err !=nil {
			log.Printf("Error Fetching ulr %s:%v",r.Url,err)
			continue
		}
		parseResult :=r.ParseFunc(body)
		requests=append(requests,parseResult.Requests...)
		for _,item :=range parseResult.Items{
			log.Printf("Got %v",item)
		}
	}
}