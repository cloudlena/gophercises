package urlshort

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//     [
//       {
//         "path": "/some-path",
//         "url": "https://www.some-url.com/demo"
//       }
//     ]
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to URLs.
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var rr []Rule
	err := json.Unmarshal(jsn, &rr)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	pathsToURLs := make(map[string]string)
	for _, r := range rr {
		pathsToURLs[r.Path] = r.URL
	}

	return MapHandler(pathsToURLs, fallback), nil
}
