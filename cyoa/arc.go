package cyoa

// Arc is a chapter of the story.
type Arc struct {
	Title   string      `json:"title"`
	Story   []string    `json:"story"`
	Options []ArcOption `json:"options"`
}

// ArcOption is an option presented to the reader during an Arc.
type ArcOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
