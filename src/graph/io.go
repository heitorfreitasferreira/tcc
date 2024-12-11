package graph

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const fileExtension = ".graph"

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
func Save(graphs []Graph, folder string) error {
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	frequency := make(map[int]int)

	for _, graph := range graphs {
		charCode := 'a' + frequency[len(graph)]
		filepath := fmt.Sprintf("%s/%d%c%s", folder, len(graph), charCode, fileExtension)
		file, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer file.Close()

		data, err := encode(graph)
		if err != nil {
			return err
		}

		_, err = file.WriteString(data)
		if err != nil {
			return err
		}
		frequency[len(graph)]++
	}
	return nil
}

// Carrega os grafos a partir da pasta que foram salvos (provavelmente com a função Save acima)
func Load(folder string) ([]Graph, error) {
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
