package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/mastertinner/gophercises/cyoa"
)

func main() {
	var arcsFilePath = flag.String(
		"arcs-file",
		"gopher.json",
		"the path to the JSON file containing the arcs of the story",
	)
	flag.Parse()

	arcs, err := cyoa.ArcsFromFile(*arcsFilePath)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting arcs from file: %s", err))
	}

	arcHandler := makeArcHandler(arcs)
	http.HandleFunc("/", arcHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// makeArcHandler creates an HTTP handler to render story arcs.
func makeArcHandler(arcs map[string]cyoa.Arc) http.HandlerFunc {
	tmpl := template.Must(template.New("page").Parse(`
		<h1>{{ .Title }}</h1>

		{{ range .Story }}
			<p>{{ . }}</p>
		{{ end }}

		{{ range .Options }}
			<button onclick="location.href='/{{ .Arc }}'" type="button">{{ .Text }}</button>
		{{ end }}
	`))

	return func(w http.ResponseWriter, r *http.Request) {
		arcName := strings.TrimPrefix(r.URL.Path, "/")
		if arcName == "" {
			arcName = "intro"
		}
		arc, ok := arcs[arcName]
		if !ok {
			code := http.StatusNotFound
			http.Error(w, http.StatusText(code), code)
			return
		}
		err := tmpl.Execute(w, arc)
		if err != nil {
			code := http.StatusInternalServerError
			http.Error(w, http.StatusText(code), code)
		}
	}
}
