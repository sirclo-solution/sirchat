package models

import (
	"errors"
)

// DialogComponent is a subtype of component. It represents a dialog component.
// DialogComponent can contain blocks.
type dialogComponent struct {
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

	// Notification is toast bar information that appears
	// with the current block.
	Notification *NotificationObject `json:"notification,omitempty"`
}

// Validate performs validation to the DrawerComponent.
func (ths *dialogComponent) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MCTDialog {
		errs = append(errs, errors.New("invalid dialog component type"))
	}

	if ths.Action != nil {
		if len(ths.Action.Buttons) == 0 {
			errs = append(errs, errors.New("there are no action buttons in the component"))
		}

		var submitCount, cancelCount int
		for i := 0; i < len(ths.Action.Buttons); i++ {
			v := ths.Action.Buttons[i]

			switch v.Type {
			case MBTTSubmit:
				submitCount++
			case MBTTCancel:
				cancelCount++
			}

			if submitCount > 1 || cancelCount > 1 {
				errs = append(errs, errors.New("there should be only one submit button and one cancel button in the action buttons"))
			}
		}
		if ths.Action.OnClose != nil {
			if valid, err := ths.Action.OnClose.Prompt.Validate(); !valid {
				errs = append(errs, err...)
			}
		}
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
func NewDialog() *dialogComponent {
	var c dialogComponent
	c.Type = MCTDialog
	c.component.IComponent = &c
	return &c
}
