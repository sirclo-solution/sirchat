package models

// ButtonBlock defines a new block of type section
type ButtonBlock struct {
	Block
	Button *ButtonBlockObject `json:"button,omitempty"`
}

type ButtonBlockObject struct {
	BlockObject
	Type  string `json:"type"`
	Label string `json:"label"`
}

func (s ButtonBlock) Validate() error {
	// ButtonBlock validation implementation

	return nil
}

// NewButtonBlock returns a new instance of a section block to be rendered
func NewButtonBlock(buttonObj *ButtonBlockObject) *ButtonBlock {
	block := ButtonBlock{
		Button: buttonObj,
	}
	block.Type = MBTButton

	return &block
}
