package cyoa

import (
	"encoding/json"
	"fmt"
	"os"
)

// ArcsFromFile reas a JSON file and returns all arcs it contains.
func ArcsFromFile(filePath string) (map[string]Arc, error) {
	arcsJSON, err := os.ReadFile("gopher.json")
	if err != nil {
		return nil, fmt.Errorf("error reading arcs file: %w", err)
	}

	var arcs map[string]Arc
	err = json.Unmarshal(arcsJSON, &arcs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling arcs JSON: %w", err)
	}

	return arcs, nil
}
