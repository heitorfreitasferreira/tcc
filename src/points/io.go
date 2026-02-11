package points

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const fileExtention = ".points"

type SaveReport struct {
	Created int
	Skipped int
}

// Salva os mapas em arquivos na pasta informada
func Save(maps []Points2D, folder string) (SaveReport, error) {
	// Cria o diretório se ele não existir
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return SaveReport{}, fmt.Errorf("failed to create directory: %w", err)
	}

	report := SaveReport{}
	frequency := make(map[int]int)

	for _, points := range maps {
		charCode := 'a' + frequency[len(points)]
		filepath := fmt.Sprintf("%s/%d%c%s", folder, len(points), charCode, fileExtention)

		_, err := os.Stat(filepath)
		if err == nil {
			frequency[len(points)]++
			report.Skipped++
			continue
		}
		if !os.IsNotExist(err) {
			return report, fmt.Errorf("stat file %q: %w", filepath, err)
		}

		enc, err := encode(points)
		if err != nil {
			return report, err
		}

		err = os.WriteFile(filepath, []byte(enc), 0o666)
		if err != nil {
			return report, err
		}

		frequency[len(points)]++
		report.Created++
	}
	return report, nil
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
