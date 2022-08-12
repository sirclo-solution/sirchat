package models

// IAppendable an interface which has an AddBlocks method.
// if there is a struct that implements an appendable,
// then the struct can use the methods contained in IAppendable
type IAppendable interface {
	// AddBlock this method can be used to append the block where it is called
	AddBlocks(blocks ...IBlock)

	// GetBlocks this method can be used to get the block where it is called
	GetBlocks()
}

// `appendable` struct type is meant to be embedded to other
// components/blocks. `appendable` provides the embedding structs
// with field `Blocks` of type `[]IBlock`.
type appendable struct {
	// Blocks is the interface that will hold the blocks
	Blocks []IBlock `json:"blocks,omitempty"`
}

// AddBlocks adds blocks (`IBlock`) to the field `Blocks` in
// the embedded struct `appendable`. Use this method as the
// alternative to adding blocks conventionally, such as when
// handling structs as IAppendable.
func (ths *appendable) AddBlocks(blocks ...IBlock) {
	ths.Blocks = append(ths.Blocks, blocks...)
}

// GetBlocks returns blocks as []IBlock from the field 'Blocks'
// in the struct. Use this method as the alternative to getting
//blocks conventionally, such as when handling structs as IAppendable.
func (ths *appendable) GetBlocks() []IBlock {
	return ths.Blocks
}
