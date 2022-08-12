package models

import "errors"

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
	GroupID     string                    `json:"group_id"`
}

// InputBlockOptionsObject is the options for radio InputBlockObject type.
type InputBlockOptionsObject struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// Validate performs validation to the ContainerBlock. Input of type
// radio should have field 'Options' defined.
func (s InputBlock) Validate() (bool, []error) {
	var errs []error
	if s.Type != MBTInput {
		errs = append(errs, errors.New("invalid input block type"))
	}

	if s.Input.Type == InputBlockObjectTypeRadio && len(s.Input.Options) == 0 {
		errs = append(errs, errors.New("radio input must have options"))
	}

	if s.Input.Type == InputBlockObjectTypeCounter && s.Input.GroupID == "" {
		errs = append(errs, errors.New("counter input must have group ID defined"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewInputBlock returns a new instance of a input block to be rendered
func NewInputBlock(inputType InputBlockObjectType) *InputBlock {
	var block InputBlock
	block.Type = MBTInput
	block.Input = &InputBlockObject{
		Type: inputType,
	}

	return &block
}

func NewInputBlockOptionsObject(value, label string) *InputBlockOptionsObject {
	return &InputBlockOptionsObject{
		Value: value,
		Label: label,
	}
}
