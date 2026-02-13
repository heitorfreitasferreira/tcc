package dataassets

import "embed"

// Files keeps data assets embedded in the binary.
//
//go:embed *.graph *.points results/summary results/evolution results/timing
var Files embed.FS
