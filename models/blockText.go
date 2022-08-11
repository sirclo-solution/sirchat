package models

import "errors"

type TextBlockObjectAlign string
type TextBlockObjectType string
type TextBlockObjectColor string

const (
	TextBlockObjectAlignCenter    TextBlockObjectAlign = "center"
	TextBlockObjectAlignLeft      TextBlockObjectAlign = "left"
	TextBlockObjectAlignRight     TextBlockObjectAlign = "right"
	TextBlockObjectTypeSpan       TextBlockObjectType  = "span"
	TextBlockObjectTypeParagraph  TextBlockObjectType  = "paragraph"
	TextBlockObjectTypeSubheading TextBlockObjectType  = "subheading"
	TextBlockObjectTypeFigure     TextBlockObjectType  = "figure"
	TextBlockObjectColorText      TextBlockObjectColor = "text"
	TextBlockObjectColorPrimary   TextBlockObjectColor = "primary"
	TextBlockObjectColorSecondary TextBlockObjectColor = "secondary"
	TextBlockObjectColorDanger    TextBlockObjectColor = "danger"
	TextBlockObjectColorDisabled  TextBlockObjectColor = "disabled"
)

// TextBlock is a subtype of block. It represents a text block.
type TextBlock struct {
	block
	Text *TextBlockObject `json:"text,omitempty"`
}

// TextBlockObject holds the detail of the TextBlock.
type TextBlockObject struct {
	appendable
	Body  string               `json:"body"`
	Align TextBlockObjectAlign `json:"align"`
	Type  TextBlockObjectType  `json:"type"`
	Color TextBlockObjectColor `json:"color"`
}

// Validate performs validation to the TextBlock. Field `Body`
// should not be empty.
func (s TextBlock) Validate() (bool, []error) {
	// TextBlock validation implementation
	var errs []error
	if s.Text.Body == "" {
		errs = append(errs, errors.New("body is missing"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewTextBlock returns a new instance of a section block to be rendered
func NewTextBlock(textObj *TextBlockObject) *TextBlock {
	block := TextBlock{
		Text: textObj,
	}
	block.Type = MBTText

	return &block
}
