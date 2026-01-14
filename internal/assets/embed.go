package assets

import (
	"embed"
	"io/fs"
)

//go:embed css js
var embedFS embed.FS

func GetAssetsFS() fs.FS {
	return embedFS
}
