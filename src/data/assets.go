package dataassets

import "embed"

// Files keeps data assets embedded in the binary.
// Includes graph/points files and experiment results (summary, evolution, timing).
//
//go:embed *.graph *.points results
var Files embed.FS
