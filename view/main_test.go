package view

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestViewRenderer(t *testing.T) {
	renderer, err := MakeRenderer()
	require.NoError(t, err)
	result, err := renderer.renderString("index.html", nil)
	require.NoError(t, err)
	require.Contains(t, result, "<html")
	require.Contains(t, result, "<body")
	t.Log("Rendered template", result)
}
