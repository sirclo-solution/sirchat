package models

// ButtonBody defines a new block of type section
type ButtonBody struct {
	Type   MessageBodyType   `json:"type"`
	Button *ButtonBodyObject `json:"button,omitempty"`
}

type ButtonBodyObject struct {
	Type   MessageButtonType `json:"type"`
	Label  string            `json:"label,omitempty"`
	Icon   string            `json:"icon,omitempty"`
	Action *Action           `json:"action,omitempty"`
}

// BodyType returns the type of the body
func (s ButtonBody) BodyType() MessageBodyType {
	return s.Type
}

// BodyType returns the type of the body
func (s ButtonBody) ButtonType() MessageButtonType {
	return s.Button.Type
}

// NewButtonBody returns a new instance of a section block to be rendered
func NewButtonBody(buttonObj *ButtonBodyObject) *ButtonBody {
	block := ButtonBody{
		Type:   MBDTImage,
		Button: buttonObj,
	}

	return &block
}
