package search

import (
	"chercher/utils/config"
	testUtils "chercher/utils/test"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
)

func TestMakeLocalSearcher(t *testing.T) {
	t.Cleanup(testUtils.ChdirToProjectRoot(t))
	cwd, _ := os.Getwd()
	localConfig := config.SearcherConfig{
		ID:   "local",
		Type: string(SearcherTypeLocal),
		Config: map[string]interface{}{
			"dir":         path.Join(cwd, "tests/mockData"),
			"ripgrepPath": "/opt/homebrew/bin/rg",
		},
	}
	localSearch, err := MakeLocalSearcher(localConfig)
	require.NoError(t, err)
	require.NotNil(t, localSearch)

	t.Run("test search", func(t *testing.T) {
		results, err := localSearch.Search("Lorem")
		require.NoError(t, err)
		require.Greater(t, len(results), 0)
	})

	t.Run("test empty result", func(t *testing.T) {
		results, err := localSearch.Search("totally doesn't exist")
		require.NoError(t, err)
		require.Equal(t, 0, len(results))
	})
}
