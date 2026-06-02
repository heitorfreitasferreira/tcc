package repository

import (
	"cmp"
	"context"
	"encoding/json"
	"io/fs"
	"path"
	"slices"
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

	slices.SortFunc(ids, func(left, right string) int {
		leftNum, leftSuffix, leftOK := splitMapID(left)
		rightNum, rightSuffix, rightOK := splitMapID(right)

		if leftOK && rightOK {
			if order := cmp.Compare(leftNum, rightNum); order != 0 {
				return order
			}

			if order := cmp.Compare(leftSuffix, rightSuffix); order != 0 {
				return order
			}

			return cmp.Compare(left, right)
		}

		if leftOK {
			return -1
		}

		if rightOK {
			return 1
		}

		return cmp.Compare(left, right)
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
