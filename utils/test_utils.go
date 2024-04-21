package utils

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func ParseHTMLToGoqueryDoc(response *http.Response) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func NodeText(node *goquery.Selection) string {
	actual := node.Text()
	return strings.Join(strings.Fields(actual), " ")
}
