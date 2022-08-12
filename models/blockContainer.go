package models

import (
	"errors"
	"fmt"
)

// ContainerBlock is a subtype of block. It represents a container block and holds
// a ContainerBlockObject in the field `Container`.
type ContainerBlock struct {
	block
	Container *ContainerBlockObject `json:"container"`
}

// ContainerBlockObject holds the detail of the ContainerBlock. ContainerBlockObject
// can contain other blocks.
type ContainerBlockObject struct {
	appendable
	Direction string `json:"direction"`
}

// Validate performs validation to the ContainerBlock.
func (s ContainerBlock) Validate() (bool, []error) {
	// ContainerBlock validation implementation
	var errs []error
	if s.Type != MBTContainer {
		errs = append(errs, errors.New("invalid container block type"))
	}

	switch s.Container.Direction {
	case "row": // add more available value here
		break
	default:
		errs = append(errs, fmt.Errorf("invalid container direction (%s)", s.Container.Direction))
	}

	if s.Container == nil {
		errs = append(errs, errors.New("field 'Container' in container block should not be empty"))
	}

	for _, v := range s.Container.Blocks {
		if valid, err := v.Validate(); !valid {
			errs = append(errs, err...)
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewContainerBlock returns a new instance of a container block to be rendered
func NewContainerBlock(direction string) *ContainerBlock {
	var block ContainerBlock
	block.Type = MBTContainer
	block.Container.Direction = direction

	return &block
}
