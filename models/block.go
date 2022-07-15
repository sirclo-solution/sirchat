package models

import (
	"encoding/json"
	"fmt"
)

type MessageBlockType string

const (
	MBTText      MessageBlockType = "text"
	MBTTable     MessageBlockType = "table"
	MBTContainer MessageBlockType = "container"
	MBTImage     MessageBlockType = "image"
)

type IBlock interface {
	BlockType() MessageBlockType
	AddBlock(blocks ...IBlock)
	Validate() bool
}

type Block struct {
	Type   MessageBlockType `json:"type"`
	Blocks []IBlock         `json:"blocks,omitempty"`
}

func (ths *Block) AddBlock(blocks ...IBlock) {
	ths.Blocks = append(ths.Blocks, blocks...)
}

// BlockType returns the type of the block
func (ths *Block) BlockType() MessageBlockType {
	return ths.Type
}

func Compose(block IBlock) []byte {
	fmt.Println("block.Validate():", block.Validate())

	res, _ := json.Marshal(block)
	return res
}

func NewBlock() *Block {
	return &Block{}
}
