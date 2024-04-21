package main

import (
	"chercher/search"
	"chercher/service"
	"chercher/view"
	"log"
)

func main() {
	searcher := search.MakeDemoSearcher()
	tmpl, err := view.MakeRenderer()
	if err != nil {
		log.Fatalln(err)
	}
	chercherService := service.MakeChercherService(searcher, tmpl)

	chercherService.ListenAndServe("127.0.0.1:8080")
}
