package view

import (
	"bytes"
	"chercher/search"
	"chercher/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRenderSearchPage(t *testing.T) {
	t.Run("test with sample data, no result", func(t *testing.T) {
		renderer, err := MakeRenderer()
		require.NoError(t, err)

		buf := &bytes.Buffer{}
		data := SearchPageDTO{SearchQuery: "hello"}
		err = renderer.RenderSearchPage(buf, data)
		require.NoError(t, err)

		doc, err := goquery.NewDocumentFromReader(buf)

		require.NoError(t, err) // valid HTML returned
		require.NotNil(t, doc)

		assert.Contains(t, utils.NodeText(doc.Find("#title")), "Found nothing that matches \"hello\"")
		assert.Equal(t, 0, doc.Find("#results > li").Length())
	})

	t.Run("test with sample data, 1 result", func(t *testing.T) {
		renderer, err := MakeRenderer()
		require.NoError(t, err)

		buf := &bytes.Buffer{}
		data := SearchPageDTO{
			SearchQuery: "hi",
			Results: []search.Result{
				{
					Title:      "hello",
					Context:    "hello",
					Href:       "https://example.com",
					Source:     "example",
					SourceIcon: "https://example.com/icon.png",
				},
			},
		}
		err = renderer.RenderSearchPage(buf, data)
		require.NoError(t, err)

		doc, err := goquery.NewDocumentFromReader(buf)

		require.NoError(t, err) // valid HTML returned
		require.NotNil(t, doc)

		assert.Contains(t, doc.Find("h1#title").Text(), "Search results for \"hi\"")
		assert.Equal(t, 1, doc.Find("#results > li").Length())
		assert.Equal(t, "https://example.com", doc.Find("a#result-0-href").AttrOr("href", ""))
		assert.Contains(t, utils.NodeText(doc.Find("#result-0-src")), "From example")
		assert.Equal(t, "https://example.com/icon.png", doc.Find("img#result-0-src-icon").AttrOr("src", ""))
	})
}
