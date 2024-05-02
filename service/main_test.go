package service

import (
	"chercher/search"
	"chercher/utils/test"
	"chercher/view"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStartServer(t *testing.T) {
	searcher := search.MakeDemoSearcher()
	tmpl, err := view.MakeRenderer()
	require.NoError(t, err)
	service := MakeChercherService(searcher, tmpl)

	makeRequest := func(request *http.Request) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		service.ServeHTTPRequest(response, request)
		return response
	}

	t.Run("Test call to /ping", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/ping", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Test asset fetch", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/dist/styles.css", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		require.Contains(t, response.Header().Get("Content-Type"), "text/css")
	})

	t.Run("Test call to /about", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/about", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		_, err := test.ParseHTMLToGoqueryDoc(response.Result())
		require.NoError(t, err)
	})

	t.Run("Test call to /", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/?q=hello", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		doc, err := test.ParseHTMLToGoqueryDoc(response.Result())
		require.NoError(t, err)

		titleNode := doc.Find("#title").First()
		require.NotNil(t, titleNode)
		titleNodeText := titleNode.Text()
		require.Contains(t, titleNodeText, `Search results for "hello"`)
	})

	t.Run("Test call to / with another query", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/?q=an%20tran", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		doc, err := test.ParseHTMLToGoqueryDoc(response.Result())
		require.NoError(t, err)

		titleNode := doc.Find("#title").First()
		require.NotNil(t, titleNode)
		titleNodeText := titleNode.Text()
		require.Contains(t, titleNodeText, `Search results for "an tran"`)
	})

	t.Run("Test call to / with a query that returns no results", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/?q=empty%20result", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		doc, err := test.ParseHTMLToGoqueryDoc(response.Result())
		require.NoError(t, err)

		titleNode := doc.Find("#title").First()
		require.NotNil(t, titleNode)
		require.Contains(t, test.NodeText(titleNode), `Found nothing that matches "empty result"`)
	})

	t.Run("Test call to / with no query", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/?q=", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		doc, err := test.ParseHTMLToGoqueryDoc(response.Result())
		require.NoError(t, err)

		require.Equal(t, 0, doc.Find("#title").Length())
	})
}
