package config

import (
	"chercher/search"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSearcherConfig(t *testing.T) {
	t.Run("test GetDataDir, normal case", func(t *testing.T) {
		sc := SearcherConfig{
			ID:      "ddg",
			DataDir: "/tmp/duckduckgo",
			Type:    string(search.SearcherTypeDuckDuckGo),
			Config: map[string]interface{}{
				"foo": "bar",
			},
		}

		assert.Equal(t, "/tmp/duckduckgo", sc.GetDataDir())
	})
	t.Run("test GetDataDir, defaulting case", func(t *testing.T) {
		sc := SearcherConfig{
			ID:   "ddg",
			Type: string(search.SearcherTypeDuckDuckGo),
			Config: map[string]interface{}{
				"foo": "bar",
			},
		}

		assert.Equal(t, "/tmp/ddg", sc.GetDataDir())
	})
}

func TestConfig_GetSearcherConfig(t *testing.T) {
	c := Config{
		SearcherConfigs: []SearcherConfig{
			{
				ID:      "ddg",
				DataDir: "/tmp/duckduckgo",
				Type:    string(search.SearcherTypeDuckDuckGo),
				Config: map[string]interface{}{
					"foo": "bar",
				},
			},
		},
	}
	config, err := c.GetSearcherConfig("ddg")
	require.NoError(t, err)
	require.NotNil(t, config)

}
