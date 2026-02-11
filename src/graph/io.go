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

func encode(g Graph) (string, error) {
	data, err := json.Marshal(g)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func decode(data string) (Graph, error) {
	var g Graph
	err := json.Unmarshal([]byte(data), &g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Salva os grafos em arquivos na pasta informada
func Save(graphs []Graph, folder string) (SaveReport, error) {
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return SaveReport{}, fmt.Errorf("failed to create directory: %w", err)
	}

	report := SaveReport{}
	frequency := make(map[int]int)

	for _, graph := range graphs {
		charCode := 'a' + frequency[len(graph)]
		filepath := fmt.Sprintf("%s/%d%c%s", folder, len(graph), charCode, fileExtension)

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
		frequency[len(graph)]++
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

// Carrega os grafos a partir da pasta que foram salvos (provavelmente com a função Save acima)
func LoadFromFolder(folder string) ([]Graph, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf("failed reading directory: %s", err)
	}

	var graphs []Graph
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

func LoadFromFile(file string) (Graph, error) {
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
