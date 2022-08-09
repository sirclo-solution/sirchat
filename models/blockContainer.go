package models

import (
	"errors"
	"fmt"
)

type ContainerBlock struct {
	block
	Container *ContainerBlockObject `json:"container"`
}

type ContainerBlockObject struct {
	appendable
	Direction string `json:"direction"`
}

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

// NewContainerBlock returns a new instance of a section block to be rendered
func NewContainerBlock(containerBlockObj *ContainerBlockObject) *ContainerBlock {
	block := ContainerBlock{
		Container: containerBlockObj,
	}
	block.Type = MBTContainer

	return &block
}
