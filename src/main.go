package main

import (
	"log"
	"net/http"
)

func MakeBaseServeMux() *http.ServeMux {
	mux := &http.ServeMux{}
	mux.Handle("GET /about", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("<html><body><h1>Hello</h1></body></html>"))
		w.Header().Set("Content-Type", "text/html")
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(200)
	}))
	mux.Handle("GET /ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("pong"))
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(200)
	}))

	mux.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		searchQuery := r.URL.Query().Get("q")
		if searchQuery == "" {
			_, err := w.Write([]byte("<html><h1>Search for something</h1></html>"))
			if err != nil {
				log.Println(err)
			}
		}
		_, err := w.Write([]byte("<html><h1>Search result for '" + searchQuery + "'</h1></html>"))
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(200)
	}))

	return mux
}

func MakeServer() *http.Server {
	return &http.Server{
		Addr:    "127.0.0.1:3000",
		Handler: MakeBaseServeMux(),
	}
}

func main() {
	server := MakeServer()

	log.Println("Listening at", server.Addr)
	log.Fatal(server.ListenAndServe())
}
