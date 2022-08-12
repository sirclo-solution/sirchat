package models

import "errors"

// IconButton is a subtype of button. It represents an action button has icon used in
// field `Action` in components and used in block components
type IconButton struct {
	button
}

// Validate performs validation to the IconButton.
func (ths *IconButton) Validate() (bool, []error) {
	var errs []error
	if ths.ButtonBlockObject.Type != MBTTAction {
		errs = append(errs, errors.New("invalid action button block object type"))
	}

	if ths.Action == nil {
		errs = append(errs, errors.New("action button icon must have action object"))
	}

	if ths.Action.ID == "" {
		errs = append(errs, errors.New("field `ID` in action object should not be empty"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewIconButton returns a new instance of a icon button to be rendered.
// This method has parameter struct ButtonBlockObject.
// Field icon, action is required.
// Field query only used when the next action requires a payload.
// Icon button has default color is primary and variant is outlined.
func NewIconButton(buttonObj ButtonBlockObject) *IconButton {
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

	var button IconButton
	button.Type = MBTTAction
	button.Action = buttonObj.Action
	button.Icon = buttonObj.Icon
	button.Color = color
	button.Variant = variant
	button.Query = buttonObj.Query

	return &button
}
