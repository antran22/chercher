package search

import (
	"chercher/utils/config"
	"fmt"
)

type DuckDuckGoSearcher struct{}

func (s *DuckDuckGoSearcher) Search(query string) ([]Result, error) {

}

func MakeDuckDuckGoSearcher(config config.SearcherConfig) (*DuckDuckGoSearcher, error) {
	if config.Type != SearcherTypeDuckDuckGo {
		return nil, fmt.Errorf("searcher type mismatch: %s", config.Type)
	}
	return &DuckDuckGoSearcher{}, nil
}
