package models

type Action struct {
	ID      string    `json:"id"`
	Buttons []IButton `json:"buttons,omitempty"`
}

func (ths *Action) AddButtons(buttons ...IButton) {
	ths.Buttons = append(ths.Buttons, buttons...)
}

func NewAction(ID string, buttons ...IButton) Action {
	return Action{
		ID:      ID,
		Buttons: buttons,
	}
}
