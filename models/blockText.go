package models

type TextBlockObjectAlign string
type TextBlockObjectType string

const (
	TextBlockObjectAlignCenter TextBlockObjectAlign = "center"
	TextBlockObjectAlignLeft   TextBlockObjectAlign = "left"
	TextBlockObjectAlignRight  TextBlockObjectAlign = "right"
	TextBlockObjectTypeSpan    TextBlockObjectType  = "span"
)

// TextBlock defines a new block of type section
type TextBlock struct {
	Block
	Text *TextBlockObject `json:"text,omitempty"`
}

type TextBlockObject struct {
	BlockObject
	Body  string               `json:"body"`
	Align TextBlockObjectAlign `json:"align"`
	Type  TextBlockObjectType  `json:"type"`
}

func (s TextBlock) Validate() (bool, error) {
	// TextBlock validation implementation

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
