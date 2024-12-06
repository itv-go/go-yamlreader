package yamlreader

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func Read[T any](filePath string, config *T) (*T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open YAML file: %v", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Failed to close YAML file: %v", err)
		}
	}(file)

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		log.Fatalf("Failed to decode YAML: %v", err)
		return nil, err
	}

	return config, nil
}
