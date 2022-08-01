package models

type SubmitButton struct {
	Button
	Label  string  `json:"label,omitempty"`
	Action *Action `json:"action,omitempty"`
}

func (s SubmitButton) Validate() (bool, error) {
	// SubmitButton validation implementation

	return true, nil
}

// NewSubmitButton returns a new instance of a section block to be rendered
func NewSubmitButton(label string) *SubmitButton {
	var button SubmitButton
	button.Label = label
	button.Type = MBTTSubmit

	return &button
}
