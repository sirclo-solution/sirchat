package models

import "errors"

// SubmitButton is a subtype of button. It represents a submit button used in
// field `Action` in components and used in block components
type SubmitButton struct {
	button
}

// Validate performs validation to the SubmitButton.
func (ths *SubmitButton) Validate() (bool, []error) {
	var errs []error
	if ths.ButtonBlockObject.Type != MBTTSubmit {
		errs = append(errs, errors.New("invalid submit button block object type"))
	}

	if ths.Label == "" {
		errs = append(errs, errors.New("submit button must have content of label object"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewSubmitButton returns a new instance of a submit button to be rendered
// This Button is used when there is input on the page of any type,
// when this button is clicked it will trigger to the next action.
// by bringing the payload that has been filled in the input.
// this method has parameter label and disabled.
// label is content or text of button,
// disabled is unclickable button.
// submit button has default color is primary and variant is contained
func NewSubmitButton(label string, disable bool) *SubmitButton {
	var button SubmitButton
	button.Label = label
	button.Type = MBTTSubmit
	button.Color = ButtonBlockObjectColorPrimary
	button.Variant = ButtonObjectVariantContained
	button.Disabled = disable

	return &button
}
