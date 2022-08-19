package models

import (
	"errors"
	"fmt"
)

type ContainerDirection string

const (
	// CDRow this is a row direction
	CDRow ContainerDirection = "row"

	// CDColumn this is a column direction
	CDColumn ContainerDirection = "column"
)

// ContainerBlock is a subtype of block. It represents a container block and holds
// a ContainerBlockObject in the field `Container`.
type containerBlock struct {
	block

	// Container contains the ButtonBlockObject that holds the detail of container block
	Container *ContainerBlockObject `json:"container"`
}

// ContainerBlockObject holds the detail of the ContainerBlock. ContainerBlockObject
// can contain other blocks.
type ContainerBlockObject struct {
	appendable

	// Direction defines the content direction in the container.
	Direction ContainerDirection `json:"direction"`
}

// Validate performs validation to the ContainerBlock.
func (ths *containerBlock) Validate() (bool, []error) {
	// ContainerBlock validation implementation
	var errs []error
	if ths.Type != MBTContainer {
		errs = append(errs, errors.New("invalid container block type"))
	}

	if directionValid := ths.Container.Direction.validateContainerObjectDirection(); !directionValid {
		errs = append(errs, fmt.Errorf("invalid ContainerDirection %v", ths.Container.Direction))
	}

	if ths.Container == nil {
		errs = append(errs, errors.New("field 'Container' in container block should not be empty"))
	}

	for _, v := range ths.Container.Blocks {
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
func NewContainerBlock(containerObj ContainerBlockObject) *containerBlock {
	obj := ContainerBlockObject{
		// default direction is column
		Direction: CDColumn,
	}

	var block containerBlock
	block.Type = MBTContainer

	if containerObj.Direction != "" {
		obj.Direction = containerObj.Direction
	}

	block.Container = &obj

	return &block
}

func (t ContainerDirection) validateContainerObjectDirection() bool {
	switch t {
	case CDColumn:
		return true
	case CDRow:
		return true
	default:
		return false
	}
}
