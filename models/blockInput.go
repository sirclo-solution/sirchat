package models

import (
	"errors"
	"fmt"
)

type InputBlockObjectType string

const (
	InputBlockObjectTypeText           InputBlockObjectType = "text"
	InputBlockObjectTypeRadio          InputBlockObjectType = "radio"
	InputBlockObjectTypeCounter        InputBlockObjectType = "counter"
	InputBlockObjectTypeNumber         InputBlockObjectType = "number"
	InputBlockObjectTypeSelect         InputBlockObjectType = "select"
	InputBlockObjectTypeDistrictSelect InputBlockObjectType = "district-select"
)

// InputBlock is a subtype of block. It represents an input block.
type InputBlock struct {
	block
	Input *InputBlockObject `json:"input,omitempty"`
}

// InputBlockObject holds the detail of the InputBlock.
type InputBlockObject struct {
	Type        InputBlockObjectType      `json:"type"`
	Value       string                    `json:"value"`
	Name        string                    `json:"name"`
	Placeholder string                    `json:"placeholder,omitempty"`
	Options     []InputBlockOptionsObject `json:"options,omitempty"`
	Label       string                    `json:"label,omitempty"`
	Tooltip     string                    `json:"tooltip,omitempty"`
	Required    bool                      `json:"required,omitempty"`
	GroupID     string                    `json:"group_id,omitempty"`
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

	if ths.Input.Type == InputBlockObjectTypeRadio && len(ths.Input.Options) == 0 {
		errs = append(errs, errors.New("radio input must have options"))
	}

	if ths.Input.Type == InputBlockObjectTypeCounter && ths.Input.GroupID == "" {
		errs = append(errs, errors.New("counter input must have group ID defined"))
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
