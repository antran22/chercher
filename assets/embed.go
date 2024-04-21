package assets

import (
	"embed"
	"net/http"
)

//go:embed all:dist
var Assets embed.FS

var AssetHandler = http.FileServer(http.FS(Assets))
