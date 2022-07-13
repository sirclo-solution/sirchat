package models

type Action struct {
	ID      string    `json:"id"`
	Buttons []Buttons `json:"buttons,omitempty"`
}

func NewAction(ID string, buttons ...Buttons) Action {
	return Action{
		ID:      ID,
		Buttons: buttons,
	}
}
