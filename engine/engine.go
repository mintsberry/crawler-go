package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching url %s", r.Url)
		fetch, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
			continue
		}

		parserResult := r.ParserFunc(fetch)
		requests = append(requests, parserResult.Request...)
		for _, item := range parserResult.Items {
			log.Printf("Get item: %s", item)
		}
	}
}
