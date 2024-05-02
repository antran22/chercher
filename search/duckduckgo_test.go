package search

import (
	"chercher/utils/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDuckDuckGoSearcher(t *testing.T) {
	t.Skip("skipping due to rate limit")
	ddg, err := MakeDuckDuckGoSearcher(config.SearcherConfig{
		Type: "duckduckgo",
	})
	require.NoError(t, err)
	require.NotNil(t, ddg)

	t.Run("test search", func(t *testing.T) {
		results, err := ddg.Search("hello")
		require.NoError(t, err)
		require.Greater(t, len(results), 0)
	})
}
