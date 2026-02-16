package repository

import (
	"io/fs"
	"strconv"
)

const (
	summaryDir   = "results/summary"
	evolutionDir = "results/evolution"
)

type EmbeddedDataRepository struct {
	dataFS fs.FS
}

func NewEmbeddedDataRepository(dataFS fs.FS) *EmbeddedDataRepository {
	return &EmbeddedDataRepository{dataFS: dataFS}
}

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
