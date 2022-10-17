package models

import (
	"errors"
	"fmt"
	"regexp"
)

// InputBlockObjectType is a type for field `Input` in InputBlockObject
type InputBlockObjectType string

const (
	// InputBlockObjectTypeText is the type for text input
	InputBlockObjectTypeText InputBlockObjectType = "text"

	// InputBlockObjectTypeTextArea is the type for text area input
	InputBlockObjectTypeTextArea InputBlockObjectType = "textarea"

	// InputBlockObjectTypePassword is the type for text area input
	InputBlockObjectTypePassword InputBlockObjectType = "password"

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

	// InputBlockObjectTypeDistrictSelect is the type for district_select input
	InputBlockObjectTypeDistrictSelect InputBlockObjectType = "district_select"

	// InputBlockObjectTypeEmail is the type for email input
	InputBlockObjectTypeEmail InputBlockObjectType = "email"
)

// InputBlock is a subtype of block. It represents an input block.
type inputBlock struct {
	block

	// Input contains the InputBlockObject that holds the detail of input block
	Input *InputBlockObject `json:"input,omitempty"`
}

// InputActionObject is an action that will trigger to the next action
// with the appropriate payload filled in the input
type InputActionObject struct {
	// ID is action id that will trigger the next action/command
	ID string `json:"id"`
}

// InputBlockObject holds the detail of the InputBlock.
type InputBlockObject struct {
	// Type is the input type. The available value is text, radio,
	// counter, number, select, district_select.
	// This field is required.
	Type InputBlockObjectType `json:"type"`

	// Value is the value that the input holds.
	// This field is required.
	// If type number or counter, the value should be number string
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
	// select, and district_select.
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
	Required *bool `json:"required"`

	// Disabled is the boolean that defines whether the input is disable or not
	// The default is false.
	// This field is optional.
	Disabled bool `json:"disabled"`

	// GroupID is the identifier for the multiple input that needs to
	// be aggregated. GroupID can be used to refer to all the input with
	// the same GroupID. Moreover, it can be used as the identifier for
	// an input that has dynamic field `Name`.
	// This field is optional.
	GroupID string `json:"group_id,omitempty"`

	// MinInput is minimal input for the block input.
	// MinInput only use in counter and number type, otherwise the field is made empty.
	// Default value is 0 and this field is optional.
	MinInput *int `json:"min_input,omitempty"`

	// MaxInput is maximal input for the block input.
	// MaxInput only use in counter, text, number, textarea, password type, otherwise the field is made empty.
	// Default value is unlimited (`empty field`) and this field is optional.
	MaxInput *int `json:"max_input,omitempty"`

	// NumberOnly is an input that is only filled with numbers if value `true`.
	// Default value is false and can only be used in text type.
	// this field is optional
	NumberOnly *bool `json:"number_only,omitempty"`

	// Field Action will trigger the next action when there is an adjustment in your input.
	// This field can only be used when the input has a trigger for the next action.
	// This field is optional.
	Action *InputActionObject `json:"action,omitempty"`
}

// InputBlockOptionsObject is the options for radio InputBlockObject type.
type InputBlockOptionsObject struct {
	// Value is a field/key which will be put in the payload
	// when proceeding to the next action
	Value string `json:"value"`

	// Label is a contain of option
	Label string `json:"label"`

	// Description is deprecated, instead of `Descriptions`
	// Description is a detail content or description of option.
	// Description is optional
	Description string `json:"description,omitempty"`

	// Descriptions is a detail contents or a list of descriptions of an option.
	// It contains an array of text blocks.
	// When the input is "select" type, the descriptions will not be rendered in the UI.
	// Descriptions is optional
	Descriptions []textBlock `json:"descriptions,omitempty"`

	// Disabled is the boolean that defines whether the input is disable or not.
	// This field props in the Options Input Block and the default is false.
	// This field is optional.
	Disabled bool `json:"disabled"`
}

// Validate performs validation to the ContainerBlock. Input of type
// radio should have field 'Options' defined.
func (ths *inputBlock) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MBTInput {
		errs = append(errs, errors.New("invalid input block type"))
	}

	if ths.Input.Name == "" {
		errs = append(errs, errors.New("input block field 'name' cannot be empty"))
	}

	// Validation value of Type Number or Counter should be number string.
	// But type Number default value is string "" and type Counter is string "0"
	if (ths.Input.Value != "" && ths.Input.Type == InputBlockObjectTypeNumber) || (ths.Input.Value != "" && ths.Input.Type == InputBlockObjectTypeCounter) {
		regexNumber := regexp.MustCompile(`\d+`)
		if ok := regexNumber.MatchString(ths.Input.Value); !ok {
			errs = append(errs, fmt.Errorf("input block field 'value' and type %v should be number string, not like this %v", ths.Input.Type, ths.Input.Value))
		}
	}

	if (ths.Input.Type == InputBlockObjectTypeRadio || ths.Input.Type == InputBlockObjectTypeSelect ||
		ths.Input.Type == InputBlockObjectTypeCheckbox) &&
		len(ths.Input.Options) == 0 {
		errs = append(errs, errors.New("radio, select, or checkbox input must have options"))
	}

	if typeValid := ths.Input.Type.validateInputObjectType(); !typeValid {
		errs = append(errs, fmt.Errorf("invalid InputBlockObjectType %v", ths.Input.Type))
	}

	for _, option := range ths.Input.Options {
		for _, description := range option.Descriptions {
			if valid, err := description.Validate(); !valid {
				errs = append(errs, err...)
			}
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewInputBlock returns a new instance of a input block to be rendered
func NewInputBlock(inputBlockObj *InputBlockObject) *inputBlock {
	var block inputBlock

	requiredTrue := true
	block.Type = MBTInput
	inputBlock := InputBlockObject{
		Type:        inputBlockObj.Type,
		Name:        inputBlockObj.Name,
		Placeholder: inputBlockObj.Placeholder,
		Label:       inputBlockObj.Label,
		Tooltip:     inputBlockObj.Tooltip,
		GroupID:     inputBlockObj.GroupID,
		Action:      inputBlockObj.Action,
		Value:       "",            // default
		Required:    &requiredTrue, // default
		Disabled:    false,         // default
	}

	if inputBlockObj.Required != nil {
		inputBlock.Required = inputBlockObj.Required
	}

	if inputBlockObj.Disabled {
		inputBlock.Disabled = inputBlockObj.Disabled
	}

	// specifically for the number type, if the value is not filled, then the default is "0"
	if inputBlockObj.Type == InputBlockObjectTypeNumber {
		inputBlock.Value = "0"
	}

	if inputBlockObj.Value != "" {
		inputBlock.Value = inputBlockObj.Value
	}

	minInputNumber := 0
	if inputBlockObj.Type == InputBlockObjectTypeNumber || inputBlockObj.Type == InputBlockObjectTypeCounter {
		inputBlock.MinInput = &minInputNumber
		if inputBlockObj.MinInput != nil {
			inputBlock.MinInput = inputBlockObj.MinInput
		}
	}

	if inputBlockObj.Type == InputBlockObjectTypeNumber || inputBlockObj.Type == InputBlockObjectTypeCounter || inputBlockObj.Type == InputBlockObjectTypeText ||
		inputBlockObj.Type == InputBlockObjectTypeTextArea || inputBlockObj.Type == InputBlockObjectTypePassword {
		if inputBlockObj.MaxInput != nil {
			inputBlock.MaxInput = inputBlockObj.MaxInput
		}
	}

	if inputBlockObj.Type == InputBlockObjectTypeText {
		var numberOnly bool
		inputBlock.NumberOnly = &numberOnly
		if inputBlockObj.NumberOnly != nil {
			inputBlock.NumberOnly = inputBlockObj.NumberOnly
		}
	}

	block.Input = &inputBlock

	return &block
}

// AddInputBlockOptionsObject adds options to field Options for input of type radio, checkbox, and select.
func (ths *inputBlock) AddInputBlockOptionsObject(optionObject InputBlockOptionsObject) {
	ths.Input.Options = append(ths.Input.Options, InputBlockOptionsObject{
		Value:        optionObject.Value,
		Label:        optionObject.Label,
		Descriptions: optionObject.Descriptions,
		Disabled:     optionObject.Disabled,
	})
}

// AddDescriptions adds descriptions to field Options for input type radio, checkbox, and select.
// When the input is type select, field Descriptions can still be added with text blocks, but the
// descriptions will not be rendered in the UI.
func (ths *InputBlockOptionsObject) AddDescriptions(descriptions ...textBlock) {
	ths.Descriptions = append(ths.Descriptions, descriptions...)
}

func (t InputBlockObjectType) validateInputObjectType() bool {
	switch t {
	case InputBlockObjectTypeText:
		return true
	case InputBlockObjectTypeTextArea:
		return true
	case InputBlockObjectTypePassword:
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
	case InputBlockObjectTypeEmail:
		return true
	default:
		return false
	}
}
