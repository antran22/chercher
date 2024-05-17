package search

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMakeAggregatedSearcher(t *testing.T) {
	searcher1 := MakeDemoSearcher()
	searcher2 := MakeDemoSearcher()
	searcher := MakeAggregatedSearcher([]Searcher{searcher1, searcher2})

	result, err := searcher.Search("hello")
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(result), 2)
}

func TestAggregatedSearcher_ParallelSearch(t *testing.T) {
	searcher1 := MakeDemoSearcherWithDelay(800, 1000)
	searcher2 := MakeDemoSearcherWithDelay(800, 1000)
	searcher := MakeAggregatedSearcher([]Searcher{searcher1, searcher2})

	st := time.Now()
	result, err := searcher.Search("hello")
	et := time.Now()
	testDuration := et.Sub(st)

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(result), 1)
	require.GreaterOrEqual(t, testDuration, 500*time.Millisecond)
	require.LessOrEqual(t, testDuration, 1000*time.Millisecond)
}
