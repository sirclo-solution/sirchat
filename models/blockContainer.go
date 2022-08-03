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

func (s ContainerBlock) Validate() (bool, error) {
	// ContainerBlock validation implementation
	if s.Type != MBTContainer {
		return false, errors.New("invalid container block type")
	}

	switch s.Container.Direction {
	case "row": // add more available value here
		break
	default:
		return false, errors.New("invalid container direction")
	}

	for i, v := range s.Container.Blocks {
		if valid, err := v.Validate(); !valid {
			return false, fmt.Errorf("Container.Blocks index %d: %s", i, err.Error())
		}
	}

	return true, nil
}

// NewContainerBlock returns a new instance of a section block to be rendered
func NewContainerBlock(containerBlockObj *ContainerBlockObject) *ContainerBlock {
	block := ContainerBlock{
		Container: containerBlockObj,
	}
	block.Type = MBTContainer

	return &block
}
