package engine

import (
	"crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching url %s", r.Url)
		result, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, result.Request...)
		for _, item := range result.Items {
			log.Printf("Get item: %s", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	fetch, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(fetch), nil
}
