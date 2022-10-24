package models

// Action contains multiple buttons block and action id.
// ID in the action will be triggered when a button with type submit is clicked
type Action struct {
	// Action ID will be triggered when a button with type submit is clicked
	ID string `json:"id"`

	// Buttons can contain multiple buttons
	Buttons []ButtonBlockObject `json:"buttons,omitempty"`

	// OnClose is used to action when clicked button X (closed).
	OnClose *ActionOnClose `json:"on_close,omitempty"`
}

type ActionOnClose struct {
	// OnClose have a trigger PromptBlock.
	Prompt *promptBlock `json:"prompt"`
}

// AddButtons used to append button on Action
func (ths *Action) AddButtons(buttons ...ButtonBlockObject) {
	ths.Buttons = append(ths.Buttons, buttons...)
}

// NewAction use to create new action
func NewAction(ID string) *Action {
	return &Action{
		ID: ID,
	}
}

// AddOnClose used to add an action that has a trigger prompt
// when clicking the X button.
func (ths *Action) AddOnClose(param *ActionOnClose) {
	ths.OnClose = param
}
