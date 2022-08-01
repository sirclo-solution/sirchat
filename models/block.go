package models

type MessageBlockType string

const (
	MBTText      MessageBlockType = "text"
	MBTTable     MessageBlockType = "table"
	MBTContainer MessageBlockType = "container"
	MBTImage     MessageBlockType = "image"
	MBTInput     MessageBlockType = "input"
	MBTButton    MessageBlockType = "button"
)

type IBlock interface {
	Validate() (bool, error)
}

// Block is the base struct for every other block type
type Block struct {
	Type MessageBlockType `json:"type"`
}

func NewBlocks(blocks ...IBlock) []IBlock {
	return blocks
}
