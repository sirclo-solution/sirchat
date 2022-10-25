package models

import (
	"errors"
)

// DrawerComponent is a subtype of component. It represents a drawer component.
// DrawerComponent can contain blocks.
type drawerComponent struct {
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

// Validate performs validation to the DrawerComponent. A drawer component
// should have its field `Action` defined. If there are any submit button or cancel
// button, each of them should only appear once.
func (ths *drawerComponent) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MCTDrawer {
		errs = append(errs, errors.New("invalid drawer component type"))
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

// NewDrawer used for initialization of new drawer components.
// NewDrawer returns a new instance of a drawer component to be rendered
func NewDrawer() *drawerComponent {
	var c drawerComponent
	c.Type = MCTDrawer
	c.component.IComponent = &c
	return &c
}
