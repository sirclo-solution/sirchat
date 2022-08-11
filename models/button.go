package models

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
