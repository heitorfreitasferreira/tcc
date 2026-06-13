package graph

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const fileExtension = ".graph"

type SaveReport struct {
	Created int
	Skipped int
}

func (g *Graph) UnmarshalJSON(data []byte) error {
	var flat struct {
		N    int       `json:"n"`
		Data []float64 `json:"data"`
	}
	if err := json.Unmarshal(data, &flat); err == nil && flat.Data != nil {
		g.N = flat.N
		g.Data = flat.Data
		return nil
	}

	var old [][][]float64
	if err := json.Unmarshal(data, &old); err != nil {
		return fmt.Errorf("graph: unrecognized format: %w", err)
	}
	n := len(old)
	g.N = n
	g.Data = make([]float64, n*n*n)
	for i := range n {
		for j := range n {
			for k := range n {
				g.Data[i*n*n+j*n+k] = old[i][j][k]
			}
		}
	}
	return nil
}

func (g Graph) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		N    int       `json:"n"`
		Data []float64 `json:"data"`
	}{
		N:    g.N,
		Data: g.Data,
	})
}

func encode(g *Graph) (string, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func decode(data string) (*Graph, error) {
	var g Graph
	err := json.Unmarshal([]byte(data), &g)
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func Save(graphs []*Graph, folder string) (SaveReport, error) {
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return SaveReport{}, fmt.Errorf("failed to create directory: %w", err)
	}

	report := SaveReport{}
	frequency := make(map[int]int)

	for _, graph := range graphs {
		charCode := 'a' + frequency[graph.N]
		filepath := fmt.Sprintf("%s/%d%c%s", folder, graph.N, charCode, fileExtension)

		data, err := encode(graph)
		if err != nil {
			return report, err
		}

		created, err := writeFileIfNotExists(filepath, []byte(data))
		if err != nil {
			return report, err
		}

		if created {
			report.Created++
		} else {
			report.Skipped++
		}
		frequency[graph.N]++
	}
	return report, nil
}

func writeFileIfNotExists(path string, data []byte) (bool, error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o666)
	if err != nil {
		if os.IsExist(err) {
			return false, nil
		}
		return false, err
	}

	if _, err := file.Write(data); err != nil {
		_ = file.Close()
		_ = os.Remove(path)
		return false, err
	}

	if err := file.Sync(); err != nil {
		_ = file.Close()
		_ = os.Remove(path)
		return false, err
	}

	if err := file.Close(); err != nil {
		_ = os.Remove(path)
		return false, err
	}

	return true, nil
}

func LoadFromFolder(folder string) ([]*Graph, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf("failed reading directory: %s", err)
	}

	var graphs []*Graph
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != fileExtension {
			continue
		}

		filepath := fmt.Sprintf("%s/%s", folder, file.Name())
		data, err := os.ReadFile(filepath)
		if err != nil {
			return nil, err
		}

		graph, err := decode(string(data))
		if err != nil {
			return nil, err
		}

		graphs = append(graphs, graph)
	}

	return graphs, nil
}

func LoadFromFile(file string) (*Graph, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	graph, err := decode(string(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode graph: %w", err)
	}

	return graph, nil
}
