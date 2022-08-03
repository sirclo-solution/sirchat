package models

type MessageButtonType string

const (
	MBTTAction MessageButtonType = "action"
	MBTTCancel MessageButtonType = "cancel"
	MBTTSubmit MessageButtonType = "submit"
)

type IButton interface {
	GetType() MessageButtonType
	Validate() (bool, error)
}

type Button struct {
	Type   MessageButtonType `json:"type"`
	Label  string            `json:"label,omitempty"`
	Action *Action           `json:"action,omitempty"`
}

func (ths *Button) GetType() MessageButtonType {
	return ths.Type
}

func NewButtons(buttons ...IButton) []IButton {
	return buttons
}
