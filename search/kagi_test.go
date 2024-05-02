package search

import (
	"chercher/utils/config"
	testUtils "chercher/utils/test"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKagiSearcher_Search(t *testing.T) {
	t.Cleanup(testUtils.ChdirToProjectRoot(t))
	kagiConfig, err := config.AppConfig.GetSearcherConfig("kagi")
	require.NoError(t, err)
	kagiSearch, err := MakeKagiSearcher(*kagiConfig)
	require.NoError(t, err)
	require.NotNil(t, kagiSearch)

	t.Run("test search", func(t *testing.T) {
		results, err := kagiSearch.Search("hello")
		require.NoError(t, err)
		require.Greater(t, len(results), 0)
	})
}
