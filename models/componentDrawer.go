package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type DrawerComponent struct {
	Component
	Title      Title      `json:"title"`
	Action     Action     `json:"action"`
	Blocks     []IBlock   `json:"blocks"`
	Subheading Subheading `json:"subheading,omitempty"`
}

func (ths *DrawerComponent) Validate() (bool, error) {
	if ths.Type != MCTDrawer {
		return false, errors.New("invalid drawer component type")
	}

	if len(ths.Action.Buttons) == 0 {
		return false, errors.New("there are no action buttons in the component")
	}

	var submitCount, cancelCount int
	for i := 1; i < len(ths.Action.Buttons); i++ {
		v := ths.Action.Buttons[i]

		switch v.GetType() {
		case MBTTSubmit:
			submitCount++
		case MBTTCancel:
			cancelCount++
		}

		if submitCount > 1 || cancelCount > 1 {
			return false, errors.New("there should be only one submit button and one cancel button in the action buttons")
		}

		if i == len(ths.Action.Buttons)-1 && (submitCount == 0 || cancelCount == 0) {
			return false, errors.New("there should be one submit button and one cancel button in the action buttons")
		}
	}

	return true, nil
}

func (ths *DrawerComponent) Compose() ([]byte, []error) {
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

func (ths *DrawerComponent) Send() (interface{}, error) {
	/*replace _ with a variable. e.x. jsonStr*/
	result, errs := ths.Compose()
	// if len(errs) != 0 {
	// 	return nil, fmt.Errorf("client Send(): %+q", errs)
	// }

	if errs != nil {
		fmt.Printf("%+q\n", errs)
		return nil, errors.New("error Blocks")
	}

	// send jsonStr to BE via http server

	return string(result), nil
}

func NewDrawer() *DrawerComponent {
	var c DrawerComponent
	c.Type = MCTDrawer
	return &c
}
