package models

import "errors"

// ButtonBlockObjectType this is type of the button
type ButtonBlockObjectType string

// ButtonBlockObjectColor this is a color of the button
type ButtonBlockObjectColor string

// ButtonBlockObjectVariant this is a variant of the button
type ButtonBlockObjectVariant string

// ButtonBlockObjectIcon this is a icon of the button
type ButtonBlockObjectIcon string

const (
	// MBTTAction this is a button type
	MBTTAction ButtonBlockObjectType = "button"

	// MBTTCancel this is a cancel type
	MBTTCancel ButtonBlockObjectType = "cancel"

	// MBTTSubmit this is a submit type
	MBTTSubmit ButtonBlockObjectType = "submit"

	// ButtonBlockObjectColorPrimary this is a primary color (#269CD9)
	ButtonBlockObjectColorPrimary ButtonBlockObjectColor = "primary"

	// ButtonBlockObjectColorSecondary this is a secondary color (#D0D9E0)
	ButtonBlockObjectColorSecondary ButtonBlockObjectColor = "secondary"

	// ButtonBlockObjectColorDanger this is a danger color (#D64241)
	ButtonBlockObjectColorDanger ButtonBlockObjectColor = "danger"

	// ButtonObjectVariantContained this is a contained variant
	ButtonObjectVariantContained ButtonBlockObjectVariant = "contained"

	// ButtonObjectVariantOutlined this is a outlined variant
	ButtonObjectVariantOutlined ButtonBlockObjectVariant = "outlined"

	// ButtonObjectVariantText this is a text variant
	ButtonObjectVariantText ButtonBlockObjectVariant = "text"

	// ButtonObjectIconCart this is a cart icon
	ButtonObjectIconCart ButtonBlockObjectIcon = "cart"

	// ButtonObjectIconView this is a view icon
	ButtonObjectIconView ButtonBlockObjectIcon = "view"

	// ButtonObjectIconTrash this is a trash icon
	ButtonObjectIconTrash ButtonBlockObjectIcon = "trash"

	// ButtonObjectIconTrash this is a delete icon (X)
	ButtonObjectIconDelete ButtonBlockObjectIcon = "delete"
)

type ButtonBlockObject struct {
	// Type is a type of button (button/action, cancel, submit).
	// This field is required.
	Type ButtonBlockObjectType `json:"type"`

	// Label is a content text of the button.
	// This field is required.
	Label string `json:"label"`

	// Color is a color of the button (primary, secondary, danger).
	// Default to primary.
	Color ButtonBlockObjectColor `json:"color,omitempty"`

	// Variant is a variant of the button (text, outlined, contained).
	// Default to outlined.
	// Ref: https://mui.com/material-ui/react-button/#main-content
	Variant ButtonBlockObjectVariant `json:"variant,omitempty"`

	// Icon is a icon on button.
	// This field can only be used for button icon.
	// This field is optional.
	Icon ButtonBlockObjectIcon `json:"icon,omitempty"`

	// Field Action will trigger the next action when the button is clicked.
	// This field can only be used when the button has a trigger for the next action.
	// This field is required only for button type "button" (action).
	Action *ButtonActionObject `json:"action,omitempty"`

	// field Query contains the payload that will be brought when the button is clicked
	// This field can only be used when the button has a trigger and a payload for the next action
	// This field is optional.
	Query interface{} `json:"query,omitempty"`

	// disabled is unclickable button,
	// default is false
	Disabled bool `json:"disabled"`
}

// ButtonActionObject
type ButtonActionObject struct {
	// ID is action id that will trigger the next action/command
	ID string `json:"id"`
}

// ButtonBlock is a subtype of block. It represents a button container block and holds
// a ButtonBlockObject in the field `Button`.
type ButtonBlock struct {
	block

	// Button contains the ButtonBlockObject that holds the detail of button block
	Button ButtonBlockObject `json:"button"`
}

// Validate performs validation to the ButtonBlock.
func (ths *ButtonBlock) Validate() (bool, []error) {
	var valid bool
	var errs []error
	switch ths.Button.Type {
	case MBTTAction:
		validateActionButton(*ths)
	case MBTTSubmit:
		validateSubmitButton(*ths)
	case MBTTCancel:
		validateCancelButton(*ths)
	}

	return valid, errs
}

// NewButtonBlock returns a new instance of a button block to be rendered
func NewButtonBlock(buttonObj *ButtonBlockObject) *ButtonBlock {
	obj := ButtonBlockObject{
		Type:     buttonObj.Type,
		Label:    buttonObj.Label,
		Color:    ButtonBlockObjectColorPrimary, // default
		Variant:  ButtonObjectVariantOutlined,   // default
		Icon:     buttonObj.Icon,
		Action:   buttonObj.Action,
		Query:    buttonObj.Query,
		Disabled: false, // default
	}

	if buttonObj.Color != "" {
		obj.Color = buttonObj.Color
	}

	if buttonObj.Variant != "" {
		obj.Variant = buttonObj.Variant
	}

	if buttonObj.Disabled {
		obj.Disabled = buttonObj.Disabled
	}

	return &ButtonBlock{block: block{Type: MBTButton}, Button: obj}
}

func validateActionButton(btn ButtonBlock) (bool, []error) {
	var errs []error
	if btn.Button.Type != MBTTAction {
		errs = append(errs, errors.New("invalid action button block object type"))
	}

	if btn.Button.Label == "" {
		errs = append(errs, errors.New("action button must have content of label object"))
	}

	if btn.Button.Action == nil {
		errs = append(errs, errors.New("action button must have action object"))
	}

	if btn.Button.Action.ID == "" {
		errs = append(errs, errors.New("field `ID` in action object should not be empty"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func validateSubmitButton(btn ButtonBlock) (bool, []error) {
	var errs []error
	if btn.Button.Type != MBTTSubmit {
		errs = append(errs, errors.New("invalid submit button block object type"))
	}

	if btn.Button.Label == "" {
		errs = append(errs, errors.New("submit button must have content of label object"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func validateCancelButton(btn ButtonBlock) (bool, []error) {
	var errs []error
	if btn.Button.Type != MBTTCancel {
		errs = append(errs, errors.New("invalid cancel button block object type"))
	}

	if btn.Button.Label == "" {
		errs = append(errs, errors.New("cancel button must have content of label object"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}
