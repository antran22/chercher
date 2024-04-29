package main

import (
	"chercher/search"
	"chercher/service"
	"chercher/utils/config"
	"chercher/view"
	"log"
	"net/http"
)

func main() {

	var searchers []search.Searcher

	for _, sc := range config.AppConfig.SearcherConfigs {
		if sc.Type == string(search.SearcherTypeDuckDuckGo) {
			ddg, err := search.MakeDuckDuckGoSearcher(sc)
			if err != nil {
				log.Println("failed to make DuckDuckGoSearcher", err)
			}
			searchers = append(searchers, ddg)
		} else if sc.Type == string(search.SearcherTypeKagi) {
			kagi, err := search.MakeKagiSearcher(sc)
			if err != nil {
				log.Println("failed to make DuckDuckGoSearcher", err)
			}
			searchers = append(searchers, kagi)
		}
	}

	combinedSearchers := search.MakeCombinedSearcher(searchers)

	renderer, err := view.MakeRenderer()
	if err != nil {
		log.Fatalln(err)
	}
	chercherService := service.MakeChercherService(combinedSearchers, renderer)

	log.Fatalln(http.ListenAndServe(config.AppConfig.ListenUrl, chercherService.GetHandler()))

	//if config.AppConfig.Mode == config.Dev {
	//	// Call `New()` with a list of directories to recursively watch
	//	reloader := reload.New(".")
	//
	//	// Optionally, define a callback to
	//	// invalidate any caches
	//	reloader.OnReload = func() {
	//		if err := renderer.Load(); err != nil {
	//			log.Println(err)
	//		}
	//	}
	//
	//	// Start watching
	//	handler := reloader.Handle(chercherService.GetHandler())
	//	log.Fatalln(http.ListenAndServe(config.AppConfig.ListenUrl, handler))
	//} else {
	//}

}
