package models

import (
	"errors"
	"fmt"
)

type ContainerBlock struct {
	Block
	Container *ContainerBlockObject `json:"container"`
}

type ContainerBlockObject struct {
	BlockObject
	Direction string `json:"direction"`
}

func (s ContainerBlock) Validate() error {
	// ContainerBlock validation implementation
	if s.Type != "container" {
		return errors.New("invalid container block type")
	}

	switch s.Container.Direction {
	case "row": // add more available value here
		break
	default:
		return errors.New("invalid container direction")
	}

	for i, v := range s.Container.Blocks {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("Container.Blocks index %d: %s", i, err.Error())
		}
	}

	return nil
}

// NewContainerBlock returns a new instance of a section block to be rendered
func NewContainerBlock(containerBlockObj *ContainerBlockObject) *ContainerBlock {
	block := ContainerBlock{
		Container: containerBlockObj,
	}
	block.Type = MBTContainer

	return &block
}
