package search

type SearcherType string

const (
	SearcherTypeDuckDuckGo SearcherType = "duckduckgo"
	SearcherTypeKagi       SearcherType = "kagi"
	SearcherTypeGoogle     SearcherType = "google"
	SearcherTypeBing       SearcherType = "bing"
	SearcherTypeLocal      SearcherType = "local"
)
