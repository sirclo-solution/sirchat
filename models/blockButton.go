package models

import "errors"

// ButtonBlock is a subtype of block. It represents a button block and holds
// a ButtonBlockObject in the field `Button`.
type ButtonBlock struct {
	block
	Button *ButtonBlockObject `json:"button,omitempty"`
}

// Validate performs validation to the ButtonBlock. The label of the buttons
// must be defined.
func (ths *ButtonBlock) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MBTButton {
		errs = append(errs, errors.New("invalid button block type"))
	}

	if ths.Button == nil {
		errs = append(errs, errors.New("field 'Button' in button block should not be empty"))
		return false, errs
	}

	if ths.Button.Type == "" {
		errs = append(errs, errors.New("field 'Type' in field 'Button' should not be empty"))
	}

	if ths.Button.Label == "" {
		errs = append(errs, errors.New("field 'Label' in field 'Button' should not be empty"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewButtonBlock returns a new instance of a button block to be rendered
func NewButtonBlock(buttonObj *ButtonBlockObject) *ButtonBlock {
	block := ButtonBlock{
		Button: buttonObj,
	}
	block.Type = MBTButton

	return &block
}
