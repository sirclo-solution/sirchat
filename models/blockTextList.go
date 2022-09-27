package models

import "errors"

// TextBlock is a subtype of block. It represents a text block.
// TextBlock contains more TextBlock
type TextListBlock struct {
	block

	// TextList is a field to contain textBlocks
	TextList []*textBlock `json:"text_list"`
}

// Validate performs validation to the TextListBlock. Field `Body`
// each TextBlock should not be empty.
func (ths *TextListBlock) Validate() (bool, []error) {
	// TextBlock validation implementation
	var errs []error
	for _, text := range ths.TextList {
		if text.Text.Body == "" {
			errs = append(errs, errors.New("body is missing"))
		}
	}

	if len(ths.TextList) == 0 {
		errs = append(errs, errors.New("text block in text list block should not be empty"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewTextListBlock returns a new instance of a section block to be rendered.
// to adding textblock to this Block, call AddTextBlock.
// you can add textblock more than 1
func NewTextListBlock() *TextListBlock {
	var block TextListBlock
	block.Type = MBTTextList
	return &block
}

// AddTextBlock use to be add textBlock on the text list block
func (ths *TextListBlock) AddTextBlock(textBlock *textBlock) {
	ths.TextList = append(ths.TextList, textBlock)
}
