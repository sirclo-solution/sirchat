package models

// Subheading is a part of drawer component. It represents a subheading
// in the drawer component. A subheading can have divider under it.
type Subheading struct {
	Text    string `json:"text"`
	Divider bool   `json:"divider,omitempty"`
}

// NewSubheading returns a new instance of a subheading object to be
// used in drawer component.
func NewSubheading(text string) Subheading {
	return Subheading{
		Text: text,
	}
}
