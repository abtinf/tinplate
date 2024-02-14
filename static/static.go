package static

import (
	"embed"
)

//go:embed http/*
var Http embed.FS
