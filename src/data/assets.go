package dataassets

import "embed"

//go:embed *.graph *.points results
var Files embed.FS
