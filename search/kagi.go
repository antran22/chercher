package search

import (
	"chercher/utils/config"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type kagiSearcher struct {
	client *http.Client
	token  string
}

func (s *kagiSearcher) Search(query string) ([]Result, error) {
	resp, err := s.client.Get("https://kagi.com/html/search?q=" + query)
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

	entries := doc.Find(".search-result")
	if entries == nil || entries.Length() == 0 {
		return []Result{}, fmt.Errorf("invalid results")
	}

	results := goquery.Map(entries, func(i int, selection *goquery.Selection) Result {
		titleEl := selection.Find("a.__sri_title_link")
		title := titleEl.Text()
		href := titleEl.AttrOr("href", "")
		context := selection.Find(".__sri-desc").Text()
		return Result{
			Href:       href,
			Title:      title,
			Source:     "Kagi",
			Context:    context,
			SourceIcon: "https://kagi.com/favicon.ico",
		}
	})

	return results, nil
}

func MakeKagiSearcher(config config.SearcherConfig) (Searcher, error) {
	if SearcherType(config.Type) != SearcherTypeKagi {
		return nil, fmt.Errorf("searcher type mismatch: %s", config.Type)
	}
	token, ok := config.Config["token"].(string)
	if !ok {
		return nil, fmt.Errorf("missing token")
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Jar: jar}
	_, err = client.Get("https://kagi.com/html/search?token=" + token)
	if err != nil {
		return nil, err
	}

	return &kagiSearcher{client: client, token: token}, nil
}
