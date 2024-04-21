package view

import (
	"chercher/search"
	"io"
)

type SearchPageDTO struct {
	SearchQuery string
	Results     []search.Result
}

func (r *Renderer) RenderSearchPage(w io.Writer, data SearchPageDTO) error {
	return r.render(w, "index.html", data)
}
