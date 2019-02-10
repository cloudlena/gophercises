package urlshort

import (
	"context"
	"fmt"
	"net/http"
)

// StoreHandler returns an http.HandlerFunc (which also implements http.Handler)
// that will dynamically to map any paths to their corresponding
// URL retrieved from a store. If the path is not provided in the store then the
// fallback http.Handler will be called instead.
//
// The only errors that can be returned all related to having
// invalid store data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to URLs.
func StoreHandler(store RuleStore, initialRules []Rule, fallback http.Handler) (http.HandlerFunc, error) {
	for _, r := range initialRules {
		err := store.Add(context.Background(), r)
		if err != nil {
			return nil, fmt.Errorf("error adding rule to store: %s", err)
		}
	}

	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		rule, err := store.Get(r.Context(), r.URL.Path)
		if err != nil {
			http.Error(w, fmt.Sprintf("error getting redirect rule in store: %s", err), http.StatusInternalServerError)
			return
		}
		if rule.URL != "" {
			http.Redirect(w, r, rule.URL, http.StatusPermanentRedirect)
			return
		}
		fallback.ServeHTTP(w, r)
	}
	return handlerFunc, nil
}
