package models

import "errors"

// ActionButton is a subtype of button. It represents an action button used in
// field `Action` in components and used in block components
type ActionButton struct {
	button
}

type ButtonActionObject struct {

	// ID is action id will be trigger to the next action / command
	ID string `json:"id"`
}

// Validate performs validation to the ActionButton.
func (ths *ActionButton) Validate() (bool, []error) {
	var errs []error
	if ths.ButtonBlockObject.Type != MBTTAction {
		errs = append(errs, errors.New("invalid action button block object type"))
	}

	if ths.Label == "" {
		errs = append(errs, errors.New("action button must have content of label object"))
	}

	if ths.Action == nil {
		errs = append(errs, errors.New("action button must have action object"))
	}

	if ths.Action.ID == "" {
		errs = append(errs, errors.New("field `ID` in action object should not be empty"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewActionButton returns a new instance of a action button to be rendered.
// This method has parameter struct ActionButton.
// Field label, action is required.
// Field query only used when the next action requires a payload.
// Action button has default color is primary and variant is outlined.
func NewActionButton(buttonObj ButtonBlockObject) *ActionButton {
	var (
		color   ButtonBlockObjectColor
		variant ButtonBlockObjectVariant
	)

	color = ButtonBlockObjectColorPrimary
	variant = ButtonObjectVariantOutlined

	if string(buttonObj.Color) != "" {
		color = buttonObj.Color
	}

	if string(buttonObj.Variant) != "" {
		variant = buttonObj.Variant
	}

	var button ActionButton
	button.Type = MBTTAction
	button.Action = buttonObj.Action
	button.Label = buttonObj.Label
	button.Color = color
	button.Variant = variant
	button.Query = buttonObj.Query
	button.Disabled = buttonObj.Disabled

	return &button
}
