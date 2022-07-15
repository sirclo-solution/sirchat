package models

type ContainerBlock struct {
	Block
	Container *ContainerBlockObject `json:"container"`
}

type ContainerBlockObject struct {
	Direction string   `json:"direction"`
	Blocks    []IBlock `json:"blocks,omitempty"`
}

func (s ContainerBlock) Validate() bool {
	// ContainerBlock validation implementation

	return true
}

// NewContainerBlock returns a new instance of a section block to be rendered
func NewContainerBlock(containerBlockObj *ContainerBlockObject) *ContainerBlock {
	block := ContainerBlock{
		Container: containerBlockObj,
	}
	block.Type = MBTContainer

	return &block
}
