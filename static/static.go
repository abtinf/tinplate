package static

import (
	"embed"
	"io/fs"
)

//go:embed http
var http embed.FS

var Http fs.FS

func init() {
	h, err := fs.Sub(http, "http")
	if err != nil {
		panic(err)
	}
	Http = h
}
