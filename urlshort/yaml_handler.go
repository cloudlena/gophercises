package urlshort

import (
	"fmt"
	"net/http"

	"github.com/go-yaml/yaml"
)

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to URLs.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var rr []Rule
	err := yaml.UnmarshalStrict(yml, &rr)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %w", err)
	}

	pathsToURLs := make(map[string]string)
	for _, r := range rr {
		pathsToURLs[r.Path] = r.URL
	}

	return MapHandler(pathsToURLs, fallback), nil
}
