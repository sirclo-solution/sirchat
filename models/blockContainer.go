package models

type ContainerBlock struct {
	Block
	Container *ContainerBlockObject `json:"container"`
}

type ContainerBlockObject struct {
	BlockObject
	Direction string `json:"direction"`
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
