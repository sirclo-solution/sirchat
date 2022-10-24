package models

import (
	"errors"
)

// PromptBlock is a subtype of block. It represents a prompt block and
// holds a PromptBlockObject in the field `Prompt`.
type promptBlock struct {
	// Prompt contains the PromptBlockObject that holds the detail of prompt block
	PromptBlockObject
}

// PromptBlockObject holds the detail of the PromptBlock. PromptBlockObject
// can contain other blocks.
type PromptBlockObject struct {
	appendable

	// Title is the title of the prompt.
	// This field is required.
	Title string `json:"title"`

	// cancelButton is the button used to trigger cancel the action.
	CancelButton *ButtonBlockObject `json:"cancelButton"`

	// continueButton is the button used to trigger continue the action.
	ContinueButton *ButtonBlockObject `json:"continueButton"`
}

// Validate performs validation to the PromptBlock.
func (ths *promptBlock) Validate() (bool, []error) {
	// PromptBlock validation implementation
	var errs []error
	if ths.CancelButton == nil {
		errs = append(errs, errors.New("CancelButton in prompt block object should not be empty"))
	}

	if ths.ContinueButton == nil {
		errs = append(errs, errors.New("ContinueButton in prompt block object should not be empty"))
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

// NewPromptBlock returns a new instance of a prompt block to be rendered
func NewPromptBlock(promptObj PromptBlockObject) *promptBlock {
	obj := PromptBlockObject{
		Title:          promptObj.Title,
		CancelButton:   promptObj.CancelButton,
		ContinueButton: promptObj.ContinueButton,
	}

	var block promptBlock
	block.PromptBlockObject = obj

	return &block
}
