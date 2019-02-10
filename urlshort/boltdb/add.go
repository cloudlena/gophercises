package boltdb

import (
	"context"
	"fmt"

	bolt "github.com/etcd-io/bbolt"
	"github.com/mastertinner/gophercises/urlshort"
)

// Add adds a rule to the store.
func (s *ruleStore) Add(_ context.Context, rule urlshort.Rule) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(s.rulesBucket))
		if b == nil {
			return fmt.Errorf("bucket %s doesn't exist", s.rulesBucket)
		}
		err := b.Put([]byte(rule.Path), []byte(rule.URL))
		if err != nil {
			return fmt.Errorf("error putting DB item: %s", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("error updating DB: %s", err)
	}
	return nil
}
