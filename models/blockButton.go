package models

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
	// Type is a type of button (button/action, cancel, submit)
	// its default each button
	Type ButtonBlockObjectType `json:"type"`

	// Label is a content text of the button
	Label string `json:"label"`

	// Color is a color of the button (primary, secondary, danger)
	Color ButtonBlockObjectColor `json:"color,omitempty"`

	// Variant is a variant of the button (text, outlined, contained)
	// Ref: https://mui.com/material-ui/react-button/#main-content
	Variant ButtonBlockObjectVariant `json:"variant,omitempty"`

	// Icon is a icon on button.
	// this field can only be used for button icon
	Icon ButtonBlockObjectIcon `json:"icon,omitempty"`

	// field Action will trigger to the next action when the button is clicked
	// This field can only be used when the button has a trigger for the next action
	Action *ButtonActionObject `json:"action,omitempty"`

	// field Query contains the payload that will be brought when the button is clicked
	// This field can only be used when the button has a trigger and a payload for the next action
	// The query can be empty if the next action does not need a payload
	Query interface{} `json:"query,omitempty"`

	// disabled is unclickable button,
	// default is false
	Disabled bool `json:"disabled"`
}

// IButton is the interface used only for the buttons in field
// `Action` in some IComponent object. Not to be confused with
// ButtonBlock.
type IButton interface {
	GetType() ButtonBlockObjectType
	Validate() (bool, []error)
}

// `button` is the base struct for every other button type. It is meant
// to be embedded to a button subtype. `button` provides the embedding
// structs with fields and the basic methods for a button.
type button struct {
	ButtonBlockObject
}

// GetType returns the type of the button. Use this method as the
// alternative for getting the value of field `Type` conventionally,
// such as when handling structs as IButton.
func (ths *button) GetType() ButtonBlockObjectType {
	return ths.Type
}
