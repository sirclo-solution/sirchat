package models

import (
	"errors"
	"fmt"
)

// ButtonBlockObjectType this is type of the button
type ButtonBlockObjectType string

// ButtonBlockObjectColor this is a color of the button
type ButtonBlockObjectColor string

// ButtonBlockObjectVariant this is a variant of the button
type ButtonBlockObjectVariant string

// ButtonBlockObjectIcon this is a icon of the button.
// there is example icon on file assets/icon
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

	// ButtonBlockObjectColorSecondary this is a secondary color (#52697A)
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

	// ButtonObjectIconNote this is a cart item note icon
	ButtonObjectIconEdit ButtonBlockObjectIcon = "edit"
)

// ButtonBlockObject holds the detail of the ButtonBlock.
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
	// This field is required only for button type "button" (action) and is optional
	// for button type "cancel". When cancel button that has action and query is
	// clicked, the current action and button that is rendered will be closed.
	Action *ButtonActionObject `json:"action,omitempty"`

	// field Query contains the payload that will be brought when the button is clicked
	// This field can only be used when the button has a trigger and a payload for the next action
	// This field is optional.
	Query interface{} `json:"query,omitempty"`

	// disabled is unclickable button,
	// default is false
	Disabled bool `json:"disabled"`

	// FullWidth is the width of the button, if true then the width is 100%.
	// otherwise the width follows the label of the button
	// default is false
	FullWidth bool `json:"full_width"`

	// Prompt is used to confirm before performing the next action,
	// either cancel or continue.
	Prompt *promptBlock `json:"prompt,omitempty"`
}

// ButtonActionObject
type ButtonActionObject struct {
	// ID is action id that will trigger the next action/command
	ID string `json:"id"`

	// Link is url that will be opened when button is clicked
	Link *string `json:"link,omitempty"`
}

// ButtonBlock is a subtype of block. It represents a button container block and holds
// a ButtonBlockObject in the field `Button`.
type buttonBlock struct {
	block

	// Button contains the ButtonBlockObject that holds the detail of button block
	Button ButtonBlockObject `json:"button"`
}

// Validate performs validation to the ButtonBlock.
func (ths *buttonBlock) Validate() (bool, []error) {
	var valid bool
	var errs []error
	switch ths.Button.Type {
	case MBTTAction:
		valid, errs = ths.validateActionButton()
	case MBTTSubmit:
		valid, errs = ths.validateSubmitButton()
	case MBTTCancel:
		valid, errs = ths.validateCancelButton()
	}

	if typeValid := ths.Button.Type.validateButtonObjectType(); !typeValid {
		errs = append(errs, fmt.Errorf("invalid ButtonBlockObjectType %v", ths.Button.Type))
	}

	if colorValid := ths.Button.Color.validateButtonObjectColor(); !colorValid {
		errs = append(errs, fmt.Errorf("invalid ButtonBlockObjectColor %v", ths.Button.Color))
	}

	if variantValid := ths.Button.Variant.validateButtonObjectVariant(); !variantValid {
		errs = append(errs, fmt.Errorf("invalid ButtonBlockObjectVariant %v", ths.Button.Variant))
	}

	if iconValid := ths.Button.Icon.validateButtonObjectIcon(); !iconValid {
		errs = append(errs, fmt.Errorf("invalid ButtonBlockObjectIcon %v", ths.Button.Icon))
	}

	if ths.Button.Prompt != nil {
		if valid, err := ths.Button.Prompt.Validate(); !valid {
			errs = append(errs, err...)
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return valid, errs
}

// NewButtonBlock returns a new instance of a button block to be rendered
func NewButtonBlock(buttonObj ButtonBlockObject) *buttonBlock {
	obj := ButtonBlockObject{
		Type:      buttonObj.Type,
		Label:     buttonObj.Label,
		Color:     ButtonBlockObjectColorPrimary, // default
		Variant:   ButtonObjectVariantOutlined,   // default
		Icon:      buttonObj.Icon,
		Action:    buttonObj.Action,
		Query:     buttonObj.Query,
		Disabled:  false, // default
		FullWidth: false, // default
		Prompt:    buttonObj.Prompt,
	}

	if obj.Icon != "" {
		obj.Color = ""
		obj.Variant = ""
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

	if buttonObj.FullWidth {
		obj.FullWidth = buttonObj.FullWidth
	}

	return &buttonBlock{block: block{Type: MBTButton}, Button: obj}
}

func (t *buttonBlock) validateActionButton() (bool, []error) {
	var errs []error
	if t.Button.Type != MBTTAction {
		errs = append(errs, errors.New("invalid action button block object type"))
	}

	if t.Button.Label == "" && t.Button.Icon == "" {
		errs = append(errs, errors.New("action button must have content of label or icon or both"))
	}

	if t.Button.Action == nil {
		errs = append(errs, errors.New("action button must have action object"))
	}

	if t.Button.Action.ID == "" {
		errs = append(errs, errors.New("field `ID` in action object should not be empty"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func (t *buttonBlock) validateSubmitButton() (bool, []error) {
	var errs []error
	if t.Button.Type != MBTTSubmit {
		errs = append(errs, errors.New("invalid submit button block object type"))
	}

	if t.Button.Label == "" {
		errs = append(errs, errors.New("submit button must have content of label object"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func (t *buttonBlock) validateCancelButton() (bool, []error) {
	var errs []error
	if t.Button.Type != MBTTCancel {
		errs = append(errs, errors.New("invalid cancel button block object type"))
	}

	if t.Button.Label == "" {
		errs = append(errs, errors.New("cancel button must have content of label object"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func (t ButtonBlockObjectType) validateButtonObjectType() bool {
	switch t {
	case MBTTAction:
		return true
	case MBTTCancel:
		return true
	case MBTTSubmit:
		return true
	default:
		return false
	}
}

func (t ButtonBlockObjectColor) validateButtonObjectColor() bool {
	switch t {
	case ButtonBlockObjectColorPrimary:
		return true
	case ButtonBlockObjectColorSecondary:
		return true
	case ButtonBlockObjectColorDanger:
		return true
	case "":
		return true
	default:
		return false
	}
}

func (t ButtonBlockObjectVariant) validateButtonObjectVariant() bool {
	switch t {
	case ButtonObjectVariantContained:
		return true
	case ButtonObjectVariantOutlined:
		return true
	case ButtonObjectVariantText:
		return true
	case "":
		return true
	default:
		return false
	}
}

func (t ButtonBlockObjectIcon) validateButtonObjectIcon() bool {
	switch t {
	case ButtonObjectIconCart:
		return true
	case ButtonObjectIconView:
		return true
	case ButtonObjectIconDelete:
		return true
	case ButtonObjectIconTrash:
		return true
	case ButtonObjectIconEdit:
		return true
	case "":
		return true
	default:
		return false
	}
}
