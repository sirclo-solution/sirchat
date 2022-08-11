package models

// CancelButton is a subtype of button. It represents an cancel button used in
// field `Action` in components.
type CancelButton struct {
	button
}

// Validate performs validation to the CancelButton.
func (ths *CancelButton) Validate() (bool, []error) {
	// CancelButton validation implementation

	return true, nil
}

// NewCancelButton returns a new instance of a cancel button to be rendered
func NewCancelButton(label string) *CancelButton {
	var button CancelButton
	button.Label = label
	button.Type = MBTTCancel

	return &button
}
