// Package repository provides read-only access to experiment artifacts embedded in
// the binary. It implements three reader interfaces (MapsReader, SummaryReader,
// EvolutionReader) over an fs.FS backed by Go's embed directive.
//
// Artifacts are stored under src/data/results/ and follow the naming conventions
// described in AGENTS-experiments.md. The repository layer abstracts the filesystem
// layout so that service and handler layers need only deal with map IDs, run IDs,
// and typed records.
package repository

import (
	"io/fs"
	"strconv"
)

const (
	summaryDir   = "results/summary"
	evolutionDir = "results/evolution"
)

// EmbeddedDataRepository implements MapsReader, SummaryReader, and EvolutionReader
// over an fs.FS. It is constructed with the embedded filesystem from src/data/assets.go.
type EmbeddedDataRepository struct {
	dataFS fs.FS
}

// NewEmbeddedDataRepository creates a repository backed by the provided fs.FS.
// Pass dataassets.Files (from tcc/data) to read the embedded experiment artifacts.
func NewEmbeddedDataRepository(dataFS fs.FS) *EmbeddedDataRepository {
	return &EmbeddedDataRepository{dataFS: dataFS}
}

// splitMapID splits a map ID like "10a" into its numeric prefix (10) and suffix ("a").
// Returns false if the ID doesn't follow the <number><letter> convention.
func splitMapID(mapID string) (int, string, bool) {
	if mapID == "" {
		return 0, "", false
	}

	split := 0
	for split < len(mapID) {
		r := mapID[split]
		if r < '0' || r > '9' {
			break
		}
		split++
	}

	if split == 0 || split == len(mapID) {
		return 0, "", false
	}

	number, err := strconv.Atoi(mapID[:split])
	if err != nil {
		return 0, "", false
	}

	return number, mapID[split:], true
}
