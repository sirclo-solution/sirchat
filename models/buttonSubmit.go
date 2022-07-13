package models

type SubmitButton struct {
	Type   MessageButtonType `json:"type"`
	Label  string            `json:"label,omitempty"`
	Action *Action           `json:"action,omitempty"`
}

// SubmitButton returns the type of the button
func (s *SubmitButton) ButtonType() MessageButtonType {
	return s.Type
}

// NewSubmitButton returns a new instance of a section block to be rendered
func NewSubmitButton(label string) *SubmitButton {
	button := SubmitButton{
		Type:  MBTTAction,
		Label: label,
	}

	return &button
}
