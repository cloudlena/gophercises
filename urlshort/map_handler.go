package urlshort

import (
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for path, url := range pathsToURLs {
			if r.URL.Path == path {
				http.Redirect(w, r, url, http.StatusPermanentRedirect)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}
}
