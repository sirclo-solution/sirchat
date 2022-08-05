package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type DialogComponent struct {
	Component
	Title  Title    `json:"title"`
	Action Action   `json:"action"`
	Blocks []IBlock `json:"blocks"`
}

func (ths *DialogComponent) Validate() (bool, error) {
	if ths.Type != MCTDrawer {
		return false, errors.New("invalid dialog component type")
	}

	return true, nil
}

func (ths *DialogComponent) Compose() ([]byte, []error) {
	var errs []error

	for i, v := range ths.Blocks {
		if valid, err := v.Validate(); !valid {
			err := fmt.Errorf("component.Blocks index %d: %s", i, err.Error())
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, errs
	}

	res, err := json.Marshal(ths)
	if err != nil {
		return nil, []error{errors.New("error when marshaling component")}
	}

	return res, nil
}

func NewDialog() *DialogComponent {
	var c DialogComponent
	c.Type = MCTDialog
	return &c
}
