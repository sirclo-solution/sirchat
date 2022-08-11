package models

import (
	"errors"
)

// DialogComponent is a subtype of component. It represents a dialog component.
type DialogComponent struct {
	component
	appendable
	Title  Title  `json:"title"`
	Action Action `json:"action"`
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

func NewDialog() *DialogComponent {
	var c DialogComponent
	c.Type = MCTDialog
	c.component.IComponent = &c
	return &c
}
