package models

import "errors"

// ActionButton is a subtype of button. It represents an action button used in
// field `Action` in components.
type ActionButton struct {
	button
	Action *ActionButtonActionObject `json:"action,omitempty"`
	Query  *interface{}              `json:"query,omitempty"`
}

type ActionButtonActionObject struct {
	ID string `json:"id"`
}

// Validate performs validation to the ActionButton.
func (ths *ActionButton) Validate() (bool, []error) {
	var errs []error
	if ths.ButtonBlockObject.Type != MBTTAction {
		errs = append(errs, errors.New("invalid action button block object type"))
	}

	if ths.Action == nil {
		errs = append(errs, errors.New("action button must have action object"))
		return false, errs
	}

	if ths.Action.ID == "" {
		errs = append(errs, errors.New("field `ID` in action object should not be empty"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewActionButton returns a new instance of a action button to be rendered
func NewActionButton(label, actionID string, query interface{}) *ActionButton {
	var button ActionButton
	button.Label = label
	button.Action = &ActionButtonActionObject{
		ID: actionID,
	}
	button.Type = MBTTAction
	button.Query = &query

	return &button
}
