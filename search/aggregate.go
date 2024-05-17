package search

import (
	"errors"
	"sync"
)

type AggregatedSearcher struct {
	searchers []Searcher
}

func MakeAggregatedSearcher(searchers []Searcher) *AggregatedSearcher {
	return &AggregatedSearcher{
		searchers: searchers,
	}
}

func (s *AggregatedSearcher) Search(query string) ([]Result, error) {
	// Todo: implement tf-idf ranking of results

	var wg sync.WaitGroup

	rc := make(chan []Result)
	ec := make(chan error)

	for _, searcher := range s.searchers {
		wg.Add(1)
		go func(searcher Searcher) {
			defer wg.Done()
			results, err := searcher.Search(query)
			if err != nil {
				ec <- err
				return
			}

			rc <- results

		}(searcher)
	}

	var wg2 sync.WaitGroup
	wg2.Add(2)

	var searchResult []Result
	var errs []error

	go func() {
		for resultSet := range rc {
			searchResult = append(searchResult, resultSet...)
		}
		wg2.Done()
	}()
	go func() {
		for err := range ec {
			errs = append(errs, err)
		}
		wg2.Done()
	}()

	wg.Wait()
	close(ec)
	close(rc)

	wg2.Wait()

	return searchResult, errors.Join(errs...)
}
