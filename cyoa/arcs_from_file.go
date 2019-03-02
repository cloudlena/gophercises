package cyoa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ArcsFromFile reas a JSON file and returns all arcs it contains.
func ArcsFromFile(filePath string) (map[string]Arc, error) {
	arcsJSON, err := ioutil.ReadFile("gopher.json")
	if err != nil {
		return nil, fmt.Errorf("error reading arcs file: %s", err)
	}

	var arcs map[string]Arc
	err = json.Unmarshal(arcsJSON, &arcs)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling arcs JSON: %s", err)
	}

	return arcs, nil
}
