package engine

import (
	"go-demo-reptile/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (se SimpleEngine) Run(seeds ...Request) {

	var requests []Request

	for _, r := range seeds {

		requests = append(requests, r)

	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			log.Printf("Fetcher err Url %s : %v", r.Url, err)
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("got item %v ", item)
		}

	}

}

func worker(r Request) (ParseResult, error) {

	log.Printf("Fetching Url: %s", r.Url)
	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher err Url %s : %v", r.Url, err)

		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(body)

	return parseResult, nil
}
