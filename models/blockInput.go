package models

// InputBlock defines a new block of type section
type InputBlock struct {
	Block
	Input *InputBlockObject `json:"input,omitempty"`
}

type InputBlockObject struct {
	BlockObject
	Type        string `json:"type"`
	Value       string `json:"value"`
	Name        string `json:"name"`
	Placeholder string `json:"placeholder"`
}

func (s InputBlock) Validate() error {
	// InputBlock validation implementation

	return nil
}

// NewInputBlock returns a new instance of a section block to be rendered
func NewInputBlock(inputObj *InputBlockObject) *InputBlock {
	block := InputBlock{
		Input: inputObj,
	}
	block.Type = MBTInput

	return &block
}
