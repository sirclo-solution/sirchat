package models

type ActionButton struct {
	Button
}

func (s ActionButton) Validate() (bool, error) {
	// ActionButton validation implementation

	return true, nil
}

// NewActionButton returns a new instance of a section block to be rendered
func NewActionButton(label, actionID string) *ActionButton {
	var button ActionButton
	button.Label = label
	button.Action = &Action{
		ID: actionID,
	}
	button.Type = MBTTAction

	return &button
}
