package models

type MessageButtonType string

const (
	MBTTAction MessageButtonType = "action"
	MBTTCancel MessageButtonType = "cancel"
	MBTTSubmit MessageButtonType = "submit"
)

type Button interface {
	ButtonType() MessageButtonType
}

type Buttons struct {
	ButtonSet []Button `json:"buttons,omitempty"`
}

func NewButtons(buttons ...Button) Buttons {
	return Buttons{
		ButtonSet: buttons,
	}
}
