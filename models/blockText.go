package models

// TextBlock defines a new block of type section
type TextBlock struct {
	Block
	Text *TextBlockObject `json:"text,omitempty"`
}

type TextBlockObject struct {
	BlockObject
	Type string `json:"type"`
	Body string `json:"body"`
}

func (s TextBlock) Validate() error {
	// TextBlock validation implementation

	return nil
}

// NewTextBlock returns a new instance of a section block to be rendered
func NewTextBlock(textObj *TextBlockObject) *TextBlock {
	block := TextBlock{
		Text: textObj,
	}
	block.Type = MBTText

	return &block
}
