package search

type SearcherType string

const (
	SearcherTypeDuckDuckGo SearcherType = "duckduckgo"
	SearcherTypeGoogle     SearcherType = "google"
	SearcherTypeBing       SearcherType = "bing"
	SearcherTypeFile       SearcherType = "file"
)
