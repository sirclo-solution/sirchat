package models

// Subheading is a part of drawer component. It represents a subheading
// in the drawer component. A subheading can have divider under it.
type Subheading struct {
	// text is a content of subheading
	Text string `json:"text"`

	// divider is block line for subheading
	// divider has default value false
	Divider bool `json:"divider,omitempty"`
}

// NewSubheading returns a new instance of a subheading object to be
// used in drawer component.
func NewSubheading(subheading Subheading) Subheading {
	return subheading
}
