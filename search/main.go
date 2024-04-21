package search

type Searcher interface {
	Search(query string) ([]Result, error)
}

type DemoSearcher struct{}

type Result struct {
	Href       string
	Title      string
	Context    string
	Source     string
	SourceIcon string
}

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
