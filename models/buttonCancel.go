package models

type CancelButton struct {
	Type   MessageButtonType `json:"type"`
	Label  string            `json:"label,omitempty"`
	Action *Action           `json:"action,omitempty"`
}

// CancelButton returns the type of the button
func (s *CancelButton) ButtonType() MessageButtonType {
	return s.Type
}

// NewCancelButton returns a new instance of a section block to be rendered
func NewCancelButton(label string) *CancelButton {
	button := CancelButton{
		Type:  MBTTAction,
		Label: label,
	}

	return &button
}
