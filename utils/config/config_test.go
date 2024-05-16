package config

import (
	"chercher/search"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearcherConfig(t *testing.T) {
	t.Run("test GetDataDir, normal case", func(t *testing.T) {
		sc := SearcherConfig{
			RootConfig: &Config{
				SearcherDataDir: "/tmp",
			},
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
			RootConfig: &Config{
				SearcherDataDir: "/tmp",
			},
			ID:   "ddg",
			Type: string(search.SearcherTypeDuckDuckGo),
			Config: map[string]interface{}{
				"foo": "bar",
			},
		}

		assert.Equal(t, "/tmp/ddg", sc.GetDataDir())
	})
}
