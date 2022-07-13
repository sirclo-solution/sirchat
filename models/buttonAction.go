package models

type ActionButton struct {
	Type   MessageButtonType `json:"type"`
	Label  string            `json:"label,omitempty"`
	Action *Action           `json:"action,omitempty"`
}

// ActionButton returns the type of the button
func (s *ActionButton) ButtonType() MessageButtonType {
	return s.Type
}

// NewActionButton returns a new instance of a section block to be rendered
func NewActionButton(label, actionID string) *ActionButton {
	button := ActionButton{
		Type:  MBTTAction,
		Label: label,
		Action: &Action{
			ID: actionID,
		},
	}

	return &button
}
