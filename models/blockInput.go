package models

type InputBlockObjectType string

const (
	InputBlockObjectTypeText    InputBlockObjectType = "text"
	InputBlockObjectTypeRadio   InputBlockObjectType = "radio"
	InputBlockObjectTypeCounter InputBlockObjectType = "counter"
)

// InputBlock defines a new block of type section
type InputBlock struct {
	Block
	Input *InputBlockObject `json:"input,omitempty"`
}

type InputBlockObject struct {
	BlockObject
	Type        string                    `json:"type"`
	Value       string                    `json:"value"`
	Name        string                    `json:"name"`
	Placeholder string                    `json:"placeholder,omitempty"`
	Options     []InputBlockOptionsObject `json:"options,omitempty"`
}

type InputBlockOptionsObject struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

func (s InputBlock) Validate() (bool, error) {
	// InputBlock validation implementation

	return true, nil
}

// NewInputBlock returns a new instance of a section block to be rendered
func NewInputBlock(inputObj *InputBlockObject) *InputBlock {
	block := InputBlock{
		Input: inputObj,
	}
	block.Type = MBTInput

	return &block
}
