package models

// Title is title on the component
type Title struct {
	// Text is the Title
	Text string `json:"text"`

	// Icon is image/icon url on the Title (png)
	// this field is optional and has default value (icon SIRCLO)
	// If this field is not filled, the icon is icon SIRCLO
	Icon string `json:"icon,omitempty"`
}

// NewTitle use to create new title
func NewTitle(title Title) Title {
	return title
}
