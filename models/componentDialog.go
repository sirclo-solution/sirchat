package models

import (
	"errors"
)

// DialogComponent is a subtype of component. It represents a dialog component.
// DialogComponent can contain blocks.
type DialogComponent struct {
	component
	appendable

	// Title is the title of the dialog.
	// This field is required.
	Title Title `json:"title"`

	// Action contains the Action object that will trigger another command/action.
	// This field is optional.
	Action *Action `json:"action,omitempty"`

	// Subheading is the text under the title.
	// This field is optional.
	Subheading *Subheading `json:"subheading,omitempty"`
}

// Validate performs validation to the DrawerComponent.
func (ths *DialogComponent) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MCTDialog {
		errs = append(errs, errors.New("invalid dialog component type"))
	}

	for _, v := range ths.Blocks {
		if valid, err := v.Validate(); !valid {
			errs = append(errs, err...)
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewDialog used for initialization of new dialog components
// NewDialog returns a new instance of a dialog component to be rendered
func NewDialog() *DialogComponent {
	var c DialogComponent
	c.Type = MCTDialog
	c.component.IComponent = &c
	return &c
}
