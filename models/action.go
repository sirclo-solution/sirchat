package models

// Action contains multiple buttons block and action id.
// ID in the action will be triggered when a button with type submit is clicked
type Action struct {
	// Action ID will be triggered when a button with type submit is clicked
	ID string `json:"id"`

	// Buttons can contain multiple buttons
	Buttons []IButton `json:"buttons,omitempty"`
}

// AddButtons used to append button on Action
func (ths *Action) AddButtons(buttons ...IButton) {
	ths.Buttons = append(ths.Buttons, buttons...)
}

// NewAction use to create new action
func NewAction(ID string) Action {
	return Action{
		ID: ID,
	}
}
