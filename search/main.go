package search

import (
	"math/rand/v2"
	"time"
)

type Searcher interface {
	Search(query string) ([]Result, error)
}

type Result struct {
	Href       string
	Title      string
	Context    string
	Source     string
	SourceIcon string
}

type DemoSearcher struct {
	minDelay int
	maxDelay int
}

func (s *DemoSearcher) Search(query string) ([]Result, error) {
	var searchResult []Result
	if query == "empty result" {
		return searchResult, nil
	}
	result := []Result{
		{
			Href:  "/demo?q=" + query,
			Title: "Demo Search Query " + query,
			Context: "Aliquam illum in quibusdam labore. Tempora soluta id soluta libero eius cum illo nesciunt illo. " +
				"Voluptatem voluptatem quo sapiente delectus sunt explicabo quo sed ex eligendi error mollitia ut. ",
		},
	}
	if s.maxDelay == 0 {
		return result, nil
	}

	delay := rand.IntN(s.maxDelay-s.minDelay) + s.minDelay
	select {
	case <-time.After(time.Duration(delay) * time.Millisecond):
		return result, nil
	}

}

func MakeDemoSearcher() *DemoSearcher {
	return &DemoSearcher{
		minDelay: 0,
		maxDelay: 0,
	}
}

func MakeDemoSearcherWithDelay(minDelay int, maxDelay int) *DemoSearcher {
	return &DemoSearcher{
		minDelay: minDelay,
		maxDelay: maxDelay,
	}
}
