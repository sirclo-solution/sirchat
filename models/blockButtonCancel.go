package models

import "errors"

// CancelButton is a subtype of button. It represents an cancel button used in
// field `Action` in components and used in block components
type CancelButton struct {
	button
}

// Validate performs validation to the CancelButton.
func (ths *CancelButton) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MBTTCancel {
		errs = append(errs, errors.New("invalid cancel button block object type"))
	}

	if ths.Label == "" {
		errs = append(errs, errors.New("cancel button must have content of label object"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewCancelButton returns a new instance of a cancel button to be rendered.
// This button is used to close the page.
// this method has parameter label.
// label is content or text of button.
// cancel button has default color is Secondary and variant is outlined.
func NewCancelButton(label string) *CancelButton {
	var button CancelButton
	button.Label = label
	button.Type = MBTTCancel
	button.Color = ButtonBlockObjectColorSecondary
	button.Variant = ButtonObjectVariantOutlined

	return &button
}
