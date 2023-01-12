package docs

import "embed"

//go:embed swagger.json swagger.yaml
var SwaggerFiles embed.FS
