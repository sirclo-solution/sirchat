package models

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
	Color TextBlockObjectColor `json:"coloer"`
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
