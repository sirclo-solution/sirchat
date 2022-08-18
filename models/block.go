package models

type MessageBlockType string

const (
	// MBTText is the type for text block
	MBTText MessageBlockType = "text"

	// MBTTextList is the type for text list block
	MBTTextList MessageBlockType = "text_list"

	// MBTTable is the type for table block
	MBTTable MessageBlockType = "table"

	// MBTContainer is the type for container block
	MBTContainer MessageBlockType = "container"

	// MBTImage is the type for image block
	MBTImage MessageBlockType = "image"

	// MBTInput is the type for input block
	MBTInput MessageBlockType = "input"

	// MBTButton is the type for button block
	MBTButton MessageBlockType = "button"

	// MBTDivider is the type for divider block
	MBTDivider MessageBlockType = "divider"

	// MBTCarousel is the type for carousel block
	MBTCarousel MessageBlockType = "carousel"
)

// IBlock is the interface for blocks. All type of blocks that embed `block`
// struct should satisfy this interface.
type IBlock interface {
	GetType() MessageBlockType
	Validate() (bool, []error)
}

// `block` is the base struct for every other block type. It is meant
// to be embedded to a block subtype. `block` provides the embedding
// structs with field `Type` and the basic methods for a block.
type block struct {
	// Type is the block type of a block.
	Type MessageBlockType `json:"type"`
}

// GetType returns the type of the block. Use this method as the
// alternative for getting the value of field `Type` conventionally,
// such as when handling structs as IBlock.
func (ths *block) GetType() MessageBlockType {
	return ths.Type
}
