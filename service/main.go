package service

import (
	"chercher/assets"
	"chercher/search"
	"chercher/view"
	"log"
	"net/http"
)

type ChercherService struct {
	viewRenderer *view.Renderer
	searcher     search.Searcher
	server       *http.Server
}

func (s *ChercherService) makeServeMux() *http.ServeMux {
	mux := &http.ServeMux{}
	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		searchQuery := r.URL.Query().Get("q")

		renderData := view.SearchPageDTO{
			SearchQuery: searchQuery,
		}

		if searchQuery != "" {
			results, err := s.searcher.Search(searchQuery)
			if err != nil {
				log.Println(err)
			} else {
				renderData.Results = results
			}
		}

		w.WriteHeader(http.StatusOK)
		err := s.viewRenderer.RenderSearchPage(w, renderData)
		if err != nil {
			log.Println(err)
		}
	}))

	mux.Handle("GET /about", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("<html><body><h1>Hello</h1></body></html>"))
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/html")
		if err != nil {
			log.Println(err)
		}
	}))

	mux.Handle("GET /ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte("pong"))
		if err != nil {
			log.Println(err)
		}

	}))

	mux.Handle("GET /dist/*", assets.AssetHandler)

	return mux
}

func (s *ChercherService) ListenAndServe(addr string) {
	s.server.Addr = addr
	log.Println("Listening on", s.server.Addr)
	log.Fatalln(s.server.ListenAndServe())
}

func (s *ChercherService) GetHandler() http.Handler {
	return s.server.Handler
}

func (s *ChercherService) ServeHTTPRequest(w http.ResponseWriter, r *http.Request) {
	s.server.Handler.ServeHTTP(w, r)
}

func MakeChercherService(searcher search.Searcher, tmpl *view.Renderer) *ChercherService {
	service := ChercherService{viewRenderer: tmpl, searcher: searcher}
	mux := service.makeServeMux()
	server := &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: mux,
	}
	service.server = server
	return &service
}
