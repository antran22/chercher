package search

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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

func TestDelayedSearcher(t *testing.T) {
	st := time.Now()
	searcher := MakeDemoSearcherWithDelay(800, 1200)
	result, err := searcher.Search("hello")
	et := time.Now()
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(result), 1)
	require.GreaterOrEqual(t, et.Sub(st), 500*time.Millisecond)
	require.LessOrEqual(t, et.Sub(st), 1500*time.Millisecond)
}
