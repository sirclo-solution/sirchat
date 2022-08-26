package models

import (
	"errors"
	"fmt"
)

// ContainerDirection is a block stack direction
type ContainerDirection string

// ContainerPosition is a position of container,
// the position is start, center, end.
// this type is automatically full height (100%) on the container.
// If the direction is row/horizontal, position start is starting from the left until position end is ends on the right.
// If the direction is column/vertical, position start is starting from the top until position end is ends on the bottom.
type ContainerPosition string

const (
	// CDRow this is a row/horizontal direction
	CDRow ContainerDirection = "row"

	// CDColumn this is a column/vertical direction
	CDColumn ContainerDirection = "column"

	// CPStart this is a starting position according to direction.
	// See more ContainerPosition for more detail
	CPStart ContainerPosition = "start"

	// CPCenter this is a center position.
	CPCenter ContainerPosition = "Center"

	// CPEnd this is a final position according to direction.
	CPEnd ContainerPosition = "end"
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

	// Position is position of container.
	// Position is optional, position is used when a container is made full height
	// and the position of the container can be adjusted (start, center, end)
	Position ContainerPosition `json:"position,omitempty"`
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

	if positionValid := ths.Container.Position.validateContainerObjectPosition(); !positionValid {
		errs = append(errs, fmt.Errorf("invalid ContainerPosition %v", ths.Container.Position))
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
		Position:  containerObj.Position,
	}

	var block containerBlock
	block.Type = MBTContainer

	if containerObj.Direction != "" {
		obj.Direction = containerObj.Direction
	}

	block.Container = &obj

	return &block
}

func (cd ContainerDirection) validateContainerObjectDirection() bool {
	switch cd {
	case CDColumn:
		return true
	case CDRow:
		return true
	default:
		return false
	}
}

func (cp ContainerPosition) validateContainerObjectPosition() bool {
	switch cp {
	case "":
		return true
	case CPStart:
		return true
	case CPCenter:
		return true
	case CPEnd:
		return true
	default:
		return false
	}
}
