package search

import (
	"chercher/utils/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDuckDuckGoSearcher(t *testing.T) {
	ddg, err := MakeDuckDuckGoSearcher(config.SearcherConfig{
		Type: "duckduckgo",
	})
	require.NoError(t, err)
	require.NotNil(t, ddg)
}
func TestDuckDuckGoSearcher_Search(t *testing.T) {
	if !*testExternalSearcher {
		t.Skip("skipping due to rate limit")
	}
	ddg, err := MakeDuckDuckGoSearcher(config.SearcherConfig{
		Type: "duckduckgo",
	})
	require.NoError(t, err)
	require.NotNil(t, ddg)
	results, err := ddg.Search("hello")
	require.NoError(t, err)
	require.Greater(t, len(results), 0)
}
