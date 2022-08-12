package models

import (
	"errors"
)

// DrawerComponent is a subtype of component. It represents a drawer component.
type DrawerComponent struct {
	component
	appendable
	Title      Title       `json:"title"`
	Action     Action      `json:"action"`
	Subheading *Subheading `json:"subheading,omitempty"`
}

// Validate performs validation to the DrawerComponent. A drawer component
// should have its field `Action` defined. If there are any submit button or cancel
// button, each of them should only appear once.
func (ths *DrawerComponent) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MCTDrawer {
		errs = append(errs, errors.New("invalid drawer component type"))
	}

	if len(ths.Action.Buttons) == 0 {
		errs = append(errs, errors.New("there are no action buttons in the component"))
	}

	var submitCount, cancelCount int
	for i := 0; i < len(ths.Action.Buttons); i++ {
		v := ths.Action.Buttons[i]

		switch v.GetType() {
		case MBTTSubmit:
			submitCount++
		case MBTTCancel:
			cancelCount++
		}

		if submitCount > 1 || cancelCount > 1 {
			errs = append(errs, errors.New("there should be only one submit button and one cancel button in the action buttons"))
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

// NewDrawer returns a new instance of a drawer component to be rendered
func NewDrawer() *DrawerComponent {
	var c DrawerComponent
	c.Type = MCTDrawer
	c.component.IComponent = &c
	return &c
}
