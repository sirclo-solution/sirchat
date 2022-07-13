package models

type MessageBlockType string

const (
	MBTText      MessageBlockType = "text"
	MBTTable     MessageBlockType = "table"
	MBTContainer MessageBlockType = "container"
)

type Block interface {
	BlockType() MessageBlockType
}

type Blocks struct {
	BlockSet []Block `json:"blocks,omitempty"`
}

func NewBlocks(blocks ...Block) Blocks {
	return Blocks{
		BlockSet: blocks,
	}
}
