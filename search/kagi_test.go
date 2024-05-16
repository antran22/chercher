package search

import (
	"chercher/utils/config"
	testUtils "chercher/utils/test"
	"flag"
	"github.com/stretchr/testify/require"
	"testing"
)

var testExternalSearcher = flag.Bool("external", false, "Test external searcher")

func TestMakeKagiSearcher(t *testing.T) {
	t.Cleanup(testUtils.ChdirToProjectRoot(t))
	kagiConfig, err := config.AppConfig.GetSearcherConfig("kagi")
	require.NoError(t, err)
	kagiSearch, err := MakeKagiSearcher(*kagiConfig)
	require.NoError(t, err)
	require.NotNil(t, kagiSearch)
}

func TestKagiSearcher_Search(t *testing.T) {
	if !*testExternalSearcher {
		t.Skip("skipping due to rate limit")
	}
	t.Cleanup(testUtils.ChdirToProjectRoot(t))
	kagiConfig, err := config.AppConfig.GetSearcherConfig("kagi")
	require.NoError(t, err)
	kagiSearch, err := MakeKagiSearcher(*kagiConfig)
	require.NoError(t, err)
	require.NotNil(t, kagiSearch)

	results, err := kagiSearch.Search("hello")
	require.NoError(t, err)
	require.Greater(t, len(results), 0)
}
