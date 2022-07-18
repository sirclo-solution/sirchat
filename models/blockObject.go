package models

type IBlockObject interface {
	AddBlock(blocks ...IBlock)
}

type BlockObject struct {
	Blocks []IBlock `json:"blocks,omitempty"`
}

func (ths *BlockObject) AddBlock(blocks ...IBlock) {
	ths.Blocks = append(ths.Blocks, blocks...)
}
