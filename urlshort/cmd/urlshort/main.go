package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	bolt "github.com/etcd-io/bbolt"
	"github.com/mastertinner/gophercises/urlshort"
	"github.com/mastertinner/gophercises/urlshort/boltdb"
)

func main() {
	var (
		rulesYAMLFilePath = flag.String(
			"rules-yaml-file",
			"rules.yml",
			"the path to the YAML file containing the redirect rules",
		)
		rulesJSONFilePath = flag.String(
			"rules-json-file",
			"rules.json",
			"the path to the JSON file containing the redirect rules",
		)
		rulesBoltDBFilePath = flag.String(
			"rules-boltdb-file",
			"rules.db",
			"the path to the BoltDB file containing the redirect rules",
		)
	)
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	rulesYAML, err := ioutil.ReadFile(*rulesYAMLFilePath)
	if err != nil {
		log.Fatal(fmt.Errorf("error reading rules YAML file: %s", err))
	}
	yamlHandler, err := urlshort.YAMLHandler(rulesYAML, mapHandler)
	if err != nil {
		log.Fatal(fmt.Errorf("error making YAML handler: %s", err))
	}

	rulesJSON, err := ioutil.ReadFile(*rulesJSONFilePath)
	if err != nil {
		log.Fatal(fmt.Errorf("error reading rules JSON file: %s", err))
	}
	jsonHandler, err := urlshort.JSONHandler(rulesJSON, yamlHandler)
	if err != nil {
		log.Fatal(fmt.Errorf("error making JSON handler: %s", err))
	}

	db, err := bolt.Open(*rulesBoltDBFilePath, 0600, nil)
	if err != nil {
		log.Fatal(fmt.Errorf("error opening bolt DB connection: %s", err))
	}
	defer db.Close()
	initialBoltDBRules := []urlshort.Rule{
		{
			Path: "/boltdb",
			URL:  "https://godoc.org/go.etcd.io/bbolt",
		},
	}
	store, err := boltdb.NewRuleStore(db, "rules")
	if err != nil {
		log.Fatal(fmt.Errorf("error creating BoltDB rule store: %s", err))
	}
	boltDBHandler, err := urlshort.StoreHandler(store, initialBoltDBRules, jsonHandler)
	if err != nil {
		log.Fatal(fmt.Errorf("error making BoltDB handler: %s", err))
	}

	fmt.Println("Starting the server on :8080")
	log.Fatal(http.ListenAndServe(":8080", boltDBHandler))
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
