package urlshort

import "context"

// Rule is a rule to forward a path to a specific URL.
type Rule struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// RuleStore allows to interact with rules.
type RuleStore interface {
	Get(ctx context.Context, path string) (Rule, error)
	Add(ctx context.Context, rule Rule) error
}
