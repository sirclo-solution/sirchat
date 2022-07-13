package models

// TextBody defines a new block of type section
type TextBody struct {
	Type MessageBodyType `json:"type"`
	Text *TextBodyObject `json:"text,omitempty"`
}

type TextBodyObject struct {
	Body string `json:"body"`
}

// BlockType returns the type of the block
func (s TextBody) BodyType() MessageBodyType {
	return s.Type
}

// NewTextBody returns a new instance of a section block to be rendered
func NewTextBody(textObj *TextBodyObject) *TextBody {
	block := TextBody{
		Type: MBDTText,
		Text: textObj,
	}

	return &block
}
