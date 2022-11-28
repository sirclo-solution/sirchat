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

// JustifyContentPosition used to aligns the flexible container's items when the items do not use all available space on the main-axis (horizontally)
// the available positions are start, center, end, and space_between
type JustifyContentPosition string

// AlignItemsPosition used to specifies the default alignment for items inside the flexible container
// the available positions are start, center, and end
type AlignItemsPosition string

const (
	// CDRow this is a row/horizontal direction
	CDRow ContainerDirection = "row"

	// CDColumn this is a column/vertical direction
	CDColumn ContainerDirection = "column"

	// CPStart this is a starting position according to direction.
	// See more ContainerPosition for more detail
	CPStart ContainerPosition = "start"

	// CPCenter this is a center position.
	CPCenter ContainerPosition = "center"

	// CPEnd this is a final position according to direction.
	CPEnd ContainerPosition = "end"

	// Items are positioned at the beginning of the container
	JCPStart JustifyContentPosition = "start"

	// Items are positioned in the center of the container
	JCPCenter JustifyContentPosition = "center"

	// Items are positioned at the end of the container
	JCPEnd JustifyContentPosition = "end"

	// Items will have space between them
	JCPSpaceBetween JustifyContentPosition = "space_between"

	// Items are positioned at the beginning of the container
	AIPStart AlignItemsPosition = "start"

	// Items are positioned at the center of the container
	AIPCenter AlignItemsPosition = "center"

	// Items are positioned at the end of the container
	AIPEnd AlignItemsPosition = "end"
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

	// JustifyContent is position between container items horizontally
	// JustifyContent is optional, used to set position between each container items horizontally
	JustifyContent JustifyContentPosition `json:"justify_content,omitempty"`

	// AlignItems is container items position vertically
	// AlignItems is optional, used to set container items position vertically
	AlignItems AlignItemsPosition `json:"align_items,omitempty"`
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

	if justifyContentValid := ths.Container.JustifyContent.validateContainerJustifyContentPosition(); !justifyContentValid {
		errs = append(errs, fmt.Errorf("invalid JustifyContent %v", ths.Container.JustifyContent))
	}

	if alignItems := ths.Container.AlignItems.validateContainerAlignItemsPosition(); !alignItems {
		errs = append(errs, fmt.Errorf("invalid AlignItems %v", ths.Container.AlignItems))
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
		Direction:      CDColumn,
		Position:       containerObj.Position,
		JustifyContent: containerObj.JustifyContent,
		AlignItems:     containerObj.AlignItems,
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

func (jcp JustifyContentPosition) validateContainerJustifyContentPosition() bool {
	switch jcp {
	case "", JCPStart, JCPCenter, JCPEnd, JCPSpaceBetween:
		return true
	default:
		return false
	}
}

func (aip AlignItemsPosition) validateContainerAlignItemsPosition() bool {
	switch aip {
	case "", AIPStart, AIPCenter, AIPEnd:
		return true
	default:
		return false
	}
}
