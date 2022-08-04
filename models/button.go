package models

type IButton interface {
	GetType() ButtonBlockObjectType
	Validate() (bool, error)
}

type Button struct {
	ButtonBlockObject
}

func (ths *Button) GetType() ButtonBlockObjectType {
	return ths.Type
}

func NewButtons(buttons ...IButton) []IButton {
	return buttons
}
