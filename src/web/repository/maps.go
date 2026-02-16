package repository

import (
	"context"
	"encoding/json"
	"io/fs"
	"path"
	"sort"
	"strings"

	"tcc/graph"
	"tcc/points"
)

func (r *EmbeddedDataRepository) ListMaps(_ context.Context) ([]string, error) {
	entries, err := fs.ReadDir(r.dataFS, ".")
	if err != nil {
		return nil, err
	}

	graphFiles := make(map[string]struct{})
	pointFiles := make(map[string]struct{})

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		extension := path.Ext(name)
		baseName := strings.TrimSuffix(name, extension)

		switch extension {
		case ".graph":
			graphFiles[baseName] = struct{}{}
		case ".points":
			pointFiles[baseName] = struct{}{}
		}
	}

	ids := make([]string, 0, len(graphFiles))
	for id := range graphFiles {
		if _, ok := pointFiles[id]; ok {
			ids = append(ids, id)
		}
	}

	sort.Slice(ids, func(i, j int) bool {
		leftNum, leftSuffix, leftOK := splitMapID(ids[i])
		rightNum, rightSuffix, rightOK := splitMapID(ids[j])

		if leftOK && rightOK {
			if leftNum != rightNum {
				return leftNum < rightNum
			}

			if leftSuffix != rightSuffix {
				return leftSuffix < rightSuffix
			}

			return ids[i] < ids[j]
		}

		if leftOK {
			return true
		}

		if rightOK {
			return false
		}

		return ids[i] < ids[j]
	})

	return ids, nil
}

func (r *EmbeddedDataRepository) LoadPoints(_ context.Context, mapID string) (points.Points2D, error) {
	rawPoints, err := fs.ReadFile(r.dataFS, mapID+".points")
	if err != nil {
		return nil, err
	}

	var loaded points.Points2D
	if err := json.Unmarshal(rawPoints, &loaded); err != nil {
		return nil, err
	}

	return loaded, nil
}

func (r *EmbeddedDataRepository) LoadGraph(_ context.Context, mapID string) (graph.Graph, error) {
	rawGraph, err := fs.ReadFile(r.dataFS, mapID+".graph")
	if err != nil {
		return nil, err
	}

	var loaded graph.Graph
	if err := json.Unmarshal(rawGraph, &loaded); err != nil {
		return nil, err
	}

	return loaded, nil
}
