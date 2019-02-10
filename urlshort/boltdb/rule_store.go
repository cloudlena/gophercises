package boltdb

import (
	"fmt"

	bolt "github.com/etcd-io/bbolt"
	"github.com/mastertinner/gophercises/urlshort"
)

type ruleStore struct {
	db          *bolt.DB
	rulesBucket string
}

// NewRuleStore creates a new rule store.
func NewRuleStore(db *bolt.DB, rulesBucket string) (urlshort.RuleStore, error) {
	s := &ruleStore{
		db:          db,
		rulesBucket: rulesBucket,
	}
	err := s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(s.rulesBucket))
		if err != nil {
			return fmt.Errorf("error creating %s bucket: %s", s.rulesBucket, err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error updating DB: %s", err)
	}
	return s, nil
}
