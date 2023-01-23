package static

import "embed"

//go:embed *.js *.css

var StaticFiles embed.FS
