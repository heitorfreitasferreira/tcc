package points

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const fileExtention = ".points"

// Salva os mapas em arquivos na pasta informada
func Save(maps []Points2D, folder string) error {
	// Cria o diretório se ele não existir
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	frequency := make(map[int]int)

	for _, points := range maps {
		charCode := 'a' + frequency[len(points)]
		filepath := fmt.Sprintf("%s/%d%c%s", folder, len(points), charCode, fileExtention)
		file, err := os.Create(filepath)
		if err != nil {
			return err
		}
		defer file.Close()
		enc, err := encode(points)
		if err != nil {
			return err
		}
		_, err = file.WriteString(enc)
		if err != nil {
			return err
		}
		frequency[len(points)]++
	}
	return nil
}

// Carrega os mapas a partir da pasta que foram salvos (provavelmente com a função Save acima)
func Load(folder string) ([]Points2D, error) {
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Fatalf("failed reading directory: %s", err)
	}

	var maps []Points2D
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != fileExtention {
			continue
		}

		filepath := fmt.Sprintf("%s/%s", folder, file.Name())
		data, err := os.ReadFile(filepath)
		if err != nil {
			log.Printf("Error reading file %s: %v", filepath, err)
			continue
		}

		points, err := decode(string(data))
		if err != nil {
			log.Printf("Error parsing file %s: %v", filepath, err)
			continue
		}

		maps = append(maps, points)
	}

	return maps, nil
}

func encode(pts Points2D) (string, error) {
	data, err := json.Marshal(pts)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func decode(data string) (Points2D, error) {
	var pts Points2D
	err := json.Unmarshal([]byte(data), &pts)
	if err != nil {
		return nil, err
	}
	return pts, nil
}
