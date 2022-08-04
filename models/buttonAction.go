package models

type ActionButton struct {
	Button
	Action *Action     `json:"action"`
	Query  interface{} `json:"query"`
}

func (s ActionButton) Validate() (bool, error) {
	// ActionButton validation implementation

	return true, nil
}

// NewActionButton returns a new instance of a section block to be rendered
func NewActionButton(label, actionID string, query interface{}) *ActionButton {
	var button ActionButton
	button.Label = label
	button.Action = &Action{
		ID: actionID,
	}
	button.Type = MBTTAction
	button.Query = query

	return &button
}
