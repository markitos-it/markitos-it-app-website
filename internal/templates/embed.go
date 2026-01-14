package templates

import (
	"embed"
	"io/fs"
)

//go:embed *.html
var embedFS embed.FS

func GetTemplatesFS() fs.FS {
	return embedFS
}
