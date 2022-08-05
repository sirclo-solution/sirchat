package models

// ButtonBlock defines a new block of type section
type ButtonBlock struct {
	Block
	Button *ButtonBlockObject `json:"button,omitempty"`
}

func (s ButtonBlock) Validate() (bool, error) {
	// ButtonBlock validation implementation

	return true, nil
}

// NewButtonBlock returns a new instance of a section block to be rendered
func NewButtonBlock(buttonObj *ButtonBlockObject) *ButtonBlock {
	block := ButtonBlock{
		Button: buttonObj,
	}
	block.Type = MBTButton

	return &block
}
