package models

// TextBlock defines a new block of type section
type TextBlock struct {
	Type MessageBlockType `json:"type"`
	Text *TextBlockObject `json:"text,omitempty"`
}

type TextBlockObject struct {
	Type string `json:"type"`
	Body string `json:"body"`
}

// BlockType returns the type of the block
func (s TextBlock) BlockType() MessageBlockType {
	return s.Type
}

// NewTextBlock returns a new instance of a section block to be rendered
func NewTextBlock(textObj *TextBlockObject) *TextBlock {
	block := TextBlock{
		Type: MBTText,
		Text: textObj,
	}

	return &block
}
