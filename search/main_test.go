package search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearcher(t *testing.T) {
	t.Run("test random search query", func(t *testing.T) {
		searcher := MakeDemoSearcher()
		result, err := searcher.Search("hello")
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(result), 1)
	})

	t.Run("test empty result", func(t *testing.T) {
		searcher := MakeDemoSearcher()
		result, err := searcher.Search("empty result")
		require.NoError(t, err)
		require.Equal(t, len(result), 0)
	})
}
