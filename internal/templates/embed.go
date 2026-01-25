package templates

import (
	"embed"
	"io/fs"
)

//go:embed shared/*.html shared/*.css shared/*.js
//go:embed home/*/*.html home/*/*.css home/*/*.js
//go:embed docs/*/*.html docs/*/*.css docs/*/*.js
//go:embed docs/*.md
var embedFS embed.FS

func FS() fs.FS {
	return embedFS
}
