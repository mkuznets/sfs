package swagger

import "embed"

//go:embed swagger.json swagger.yaml
var Specs embed.FS
