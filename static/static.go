package static

import (
	"embed"
	"gonfoot/internal/nolsfs"
	"io/fs"
)

//go:embed http
var http embed.FS

var Http fs.FS

func init() {
	h, err := fs.Sub(nolsfs.New(http, "index.html"), "http")
	if err != nil {
		panic(err)
	}
	Http = h
}
