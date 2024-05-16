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
		switch sc.Type {
		case string(search.SearcherTypeDuckDuckGo):
			ddg, err := search.MakeDuckDuckGoSearcher(sc)
			if err != nil {
				log.Println("failed to make DuckDuckGoSearcher:", err)
			} else {
				searchers = append(searchers, ddg)
			}
		case string(search.SearcherTypeKagi):
			kagi, err := search.MakeKagiSearcher(sc)
			if err != nil {
				log.Println("failed to make KagiSearcher:", err)
			} else {
				searchers = append(searchers, kagi)
			}
		case string(search.SearcherTypeLocal):
			local, err := search.MakeLocalSearcher(sc)
			if err != nil {
				log.Println("failed to make LocalSearcher:", err)
			} else {
				searchers = append(searchers, local)
			}
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
