package assets

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeEmbed(t *testing.T) {
	t.Run("access styles.css", func(t *testing.T) {
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/dist/styles.css", nil)
		AssetHandler.ServeHTTP(response, request)
		require.Equal(t, http.StatusOK, response.Code)
		require.Contains(t, response.Header().Get("Content-Type"), "text/css")
	})
}
