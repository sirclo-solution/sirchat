package models

type DividerBlock struct {
	Block
}

func (s DividerBlock) Validate() (bool, error) {
	// DividerBlock validation implementation

	return true, nil
}

// NewButtonBlock returns a new instance of a section block to be rendered
func NewDividerBlock() *DividerBlock {
	block := DividerBlock{}
	block.Type = MBTDivider

	return &block
}
