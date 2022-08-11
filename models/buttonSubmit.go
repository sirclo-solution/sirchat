package models

import "errors"

// SubmitButton is a subtype of button. It represents a submit button used in
// field `Action` in components.
type SubmitButton struct {
	button
	Label  string  `json:"label,omitempty"`
	Action *Action `json:"action,omitempty"`
}

// Validate performs validation to the SubmitButton.
func (ths *SubmitButton) Validate() (bool, []error) {
	var errs []error
	if ths.ButtonBlockObject.Type != MBTTSubmit {
		errs = append(errs, errors.New("invalid submit button block object type"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewSubmitButton returns a new instance of a submit button to be rendered
func NewSubmitButton(label string) *SubmitButton {
	var button SubmitButton
	button.Label = label
	button.Type = MBTTSubmit

	return &button
}
