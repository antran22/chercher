package test

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
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

func ChdirToProjectRoot(t *testing.T) func() {
	wd, err := os.Getwd()
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	t.Log("changing working directory to", dir)
	err = os.Chdir(dir)
	if err != nil {
		t.Fatalf("error changing working directory to root: %v", err)
	}
	return func() {
		if err := os.Chdir(wd); err != nil {
			t.Fatalf("error restoring working directory: %v", err)
		}
	}
}
