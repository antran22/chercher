package search

import (
	"chercher/utils/config"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type duckduckgoSearcher struct {
	client *http.Client
}

func (s *duckduckgoSearcher) Search(query string) ([]Result, error) {
	// Adapt from https://github.com/deedy5/duckduckgo_search/blob/main/duckduckgo_search/duckduckgo_search_async.py
	params := url.Values{
		"q":   {query},
		"b":   {""},
		"kl":  {"wt-wt"},
		"api": {"d.js"},
		"o":   {"json"},
		"p":   {"-2"},
	}
	req, err := http.NewRequest("POST", "https://html.duckduckgo.com/html", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:125.0) Gecko/20100101 Firefox/125.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("error closing response body", err)
		}
	}(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}

	selections := doc.Find("#links > .result")
	if selections == nil || selections.Length() == 0 {
		return []Result{}, fmt.Errorf("invalid results")
	}

	results := goquery.Map(selections, func(i int, selection *goquery.Selection) Result {
		titleEl := selection.Find("a.result__a")
		title := titleEl.Text()
		href := titleEl.AttrOr("href", "")
		context := selection.Find(".result__snippet").Text()
		return Result{
			Href:       href,
			Title:      title,
			Source:     "DuckDuckGo",
			Context:    context,
			SourceIcon: "https://duckduckgo.com/favicon.ico",
		}
	})

	return results, nil
}

func MakeDuckDuckGoSearcher(config config.SearcherConfig) (Searcher, error) {
	if SearcherType(config.Type) != SearcherTypeDuckDuckGo {
		return nil, fmt.Errorf("searcher type mismatch: %s", config.Type)
	}
	client := &http.Client{}
	return &duckduckgoSearcher{client: client}, nil
}
