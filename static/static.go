package static

import "embed"

//go:embed *.html *.js *.css

var StaticFiles embed.FS
