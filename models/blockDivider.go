package models

// DividerBlock is a subtype of block. It represents a divider block. It
// will render as a divider line between blocks.
type DividerBlock struct {
	block
}

// Validate performs validation to the DividerBlock.
func (s DividerBlock) Validate() (bool, []error) {
	// DividerBlock validation implementation

	return true, nil
}

// NewDividerBlock returns a new instance of a divider block to be rendered
func NewDividerBlock() *DividerBlock {
	block := DividerBlock{}
	block.Type = MBTDivider

	return &block
}
