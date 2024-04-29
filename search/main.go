package search

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

type DemoSearcher struct{}

func (s *DemoSearcher) Search(query string) ([]Result, error) {
	var searchResult []Result
	if query == "empty result" {
		return searchResult, nil
	}
	return []Result{
		{
			Href:  "https://duckduckgo.com/?q=" + query,
			Title: "DuckDuckGo " + query,
			Context: "Aliquam illum in quibusdam labore. Tempora soluta id soluta libero eius cum illo nesciunt illo. " +
				"Voluptatem voluptatem quo sapiente delectus sunt explicabo quo sed ex eligendi error mollitia ut. ",
		},
	}, nil
}

func MakeDemoSearcher() Searcher {
	return &DemoSearcher{}
}

type AggregatedSearcher struct {
	searchers []Searcher
}

func MakeCombinedSearcher(searchers []Searcher) Searcher {
	return &AggregatedSearcher{
		searchers: searchers,
	}
}

func (s *AggregatedSearcher) Search(query string) ([]Result, error) {
	var searchResult []Result
	for _, searcher := range s.searchers {
		results, err := searcher.Search(query)
		if err != nil {
			return nil, err
		}
		searchResult = append(searchResult, results...)
	}
	return searchResult, nil
}
