package models

import (
	"errors"
	"fmt"
)

// InputBlockObjectType is a type for field `Input` in InputBlockObject
type InputBlockObjectType string

const (
	// InputBlockObjectTypeText is the type for text input
	InputBlockObjectTypeText InputBlockObjectType = "text"

	// InputBlockObjectTypeRadio is the type for radio input
	InputBlockObjectTypeRadio InputBlockObjectType = "radio"

	// InputBlockObjectTypeCounter is the type for counter input
	InputBlockObjectTypeCounter InputBlockObjectType = "counter"

	// InputBlockObjectTypeCheckbox is the type for checkbox input
	InputBlockObjectTypeCheckbox InputBlockObjectType = "checkbox"

	// InputBlockObjectTypeNumber is the type for number input
	InputBlockObjectTypeNumber InputBlockObjectType = "number"

	// InputBlockObjectTypeSelect is the type for select input
	InputBlockObjectTypeSelect InputBlockObjectType = "select"

	// InputBlockObjectTypeDistrictSelect is the type for district-select input
	InputBlockObjectTypeDistrictSelect InputBlockObjectType = "district-select"
)

// InputBlock is a subtype of block. It represents an input block.
type InputBlock struct {
	block

	// Input contains the InputBlockObject that holds the detail of input block
	Input *InputBlockObject `json:"input,omitempty"`
}

// InputBlockObject holds the detail of the InputBlock.
type InputBlockObject struct {
	// Type is the input type. The available value is text, radio,
	// counter, number, select, district-select.
	// This field is required.
	Type InputBlockObjectType `json:"type"`

	// Value is the value that the input holds.
	// This field is required.
	Value string `json:"value"`

	// Name is the unique identifier for the input. It can be used as a
	// reference to the input.
	// This field is required.
	Name string `json:"name"`

	// Placeholder is the text that will show as the hint for user to
	// fill the input. It can only be rendered in input other than radio button
	// and checkbox.
	// This field is optional.
	Placeholder string `json:"placeholder,omitempty"`

	// Options defines the list for options in input type: radio, checkbox,
	// select, and district-select.
	// This field is optional for input besides the input mentioned above.
	Options []InputBlockOptionsObject `json:"options,omitempty"`

	// Label is the text that tells user what the input is for.
	// This field is optional.
	Label string `json:"label,omitempty"`

	// Tooltip is the text that shows the more detailed guide about
	// the input.
	// This field is optional.
	Tooltip string `json:"tooltip,omitempty"`

	// Required is the boolean that defines whether the input is need
	// to be filled or not.
	// The default is true.
	// This field is optional.
	Required bool `json:"required,omitempty"`

	// GroupID is the identifier for the multiple input that needs to
	// be aggregated. GroupID can be used to refer to all the input with
	// the same GroupID. Moreover, it can be used as the identifier for
	// an input that has dynamic field `Name`.
	// This field is optional.
	GroupID string `json:"group_id,omitempty"`
}

// InputBlockOptionsObject is the options for radio InputBlockObject type.
type InputBlockOptionsObject struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// Validate performs validation to the ContainerBlock. Input of type
// radio should have field 'Options' defined.
func (ths *InputBlock) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MBTInput {
		errs = append(errs, errors.New("invalid input block type"))
	}

	if ths.Input.Name == "" {
		errs = append(errs, errors.New("input block field 'name' cannot be empty"))
	}

	if ths.Input.Value == "" {
		errs = append(errs, errors.New("input block field 'name' cannot be empty"))
	}

	if (ths.Input.Type == InputBlockObjectTypeRadio || ths.Input.Type == InputBlockObjectTypeSelect ||
		ths.Input.Type == InputBlockObjectTypeCheckbox || ths.Input.Type == InputBlockObjectTypeCounter) &&
		len(ths.Input.Options) == 0 {
		errs = append(errs, errors.New("radio, select, district-select, checkbox, or counter input must have options"))
	}

	if typeValid := ths.Input.Type.validateInputObjectType(); !typeValid {
		errs = append(errs, fmt.Errorf("invalid InputBlockObjectType %v", ths.Input.Type))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewInputBlock returns a new instance of a input block to be rendered
func NewInputBlock(inputBlockObj *InputBlockObject) *InputBlock {
	var block InputBlock
	block.Type = MBTInput
	block.Input = inputBlockObj

	return &block
}

// AddInputBlockOptionsObject adds options to field Options for input of type "select"
func (ths *InputBlock) AddInputBlockOptionsObject(value, label string) {
	ths.Input.Options = append(ths.Input.Options, InputBlockOptionsObject{
		Value: value,
		Label: label,
	})
}

func (t InputBlockObjectType) validateInputObjectType() bool {
	switch t {
	case InputBlockObjectTypeText:
		return true
	case InputBlockObjectTypeRadio:
		return true
	case InputBlockObjectTypeCounter:
		return true
	case InputBlockObjectTypeNumber:
		return true
	case InputBlockObjectTypeSelect:
		return true
	case InputBlockObjectTypeDistrictSelect:
		return true
	default:
		return false
	}
}
