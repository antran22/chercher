package main

import (
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStartServer(t *testing.T) {
	server := MakeServer()

	makeRequest := func(request *http.Request) *httptest.ResponseRecorder {
		response := httptest.NewRecorder()
		server.Handler.ServeHTTP(response, request)
		return response
	}

	extractRequestBodyString := func(response *http.Response) (string, error) {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		return string(body), nil
	}

	t.Run("Assert server endpoint", func(t *testing.T) {
		require.Contains(t, server.Addr, ":3000")
	})

	t.Run("Test call to /ping", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/ping", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Test call to /about", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/about", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		body, err := extractRequestBodyString(response.Result())
		require.NoError(t, err)
		require.Contains(t, body, "<html")
	})

	t.Run("Test call to /", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/?q=hello", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		body, err := extractRequestBodyString(response.Result())
		require.NoError(t, err)
		require.Contains(t, body, "hello")
	})

	t.Run("Test call to / with another query", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/?q=an%20tran", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		body, err := extractRequestBodyString(response.Result())
		require.NoError(t, err)
		require.Contains(t, body, "an tran")
	})

	t.Run("Test call to / with no query", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/?q=", nil)
		response := makeRequest(request)
		require.Equal(t, http.StatusOK, response.Code)
		body, err := extractRequestBodyString(response.Result())
		require.NoError(t, err)
		require.Contains(t, body, "Search for something")
	})
}
