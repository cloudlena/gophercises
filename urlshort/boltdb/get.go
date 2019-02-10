package boltdb

import (
	"context"
	"fmt"

	bolt "github.com/etcd-io/bbolt"
	"github.com/mastertinner/gophercises/urlshort"
)

// Get retrieves all rules from the DB.
func (s *ruleStore) Get(_ context.Context, path string) (urlshort.Rule, error) {
	var rule urlshort.Rule
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.rulesBucket))
		if b == nil {
			return fmt.Errorf("bucket %s doesn't exist", s.rulesBucket)
		}
		url := b.Get([]byte(path))
		rule = urlshort.Rule{
			Path: path,
			URL:  string(url),
		}
		return nil
	})
	if err != nil {
		return urlshort.Rule{}, fmt.Errorf("error viewing DB: %s", err)
	}
	return rule, nil
}
