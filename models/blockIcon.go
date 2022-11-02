package models

import (
	"errors"
	"fmt"
	"regexp"
)

// channelIconBlock is a subtype of block. It represents an icon block.
type iconBlock struct {
	block

	// Icon contains the IconBlockObject that holds the detail of icon block
	Icon *IconBlockObject `json:"image"`
}

// ImageBlockObject holds the detail of the ImageBlock.
type IconBlockObject struct {
	// Src is the icon url
	Src string `json:"src"`

	// Alt is the alternate text or label on icon
	Alt string `json:"alt"`
}

// Validate performs validation to the iconBlock.
func (ths *iconBlock) Validate() (bool, []error) {
	// ImageBlock validation implementation
	var errs []error
	if ths.Type != MBTIcon {
		errs = append(errs, errors.New("invalid icon block type"))
	}

	match, err := regexp.MatchString(`^(https:\/\/)([^\s(["<,>/]*[^/]*)(\/)[^\s[",><]*[^/]*((\?{1}|\#{1}).*)?$`, ths.Icon.Src)
	if err != nil {
		errs = append(errs, fmt.Errorf("error when validating icon src: %s", err.Error()))
	} else if !match {
		errs = append(errs, fmt.Errorf("invalid icon src: %s", ths.Icon.Src))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewIconBlock returns a new instance of a section block to be rendered
func NewIconBlock(iconObj IconBlockObject) *iconBlock {
	block := iconBlock{}
	block.Type = MBTIcon
	block.Icon = &IconBlockObject{
		Src: iconObj.Src,
		Alt: iconObj.Alt,
	}

	return &block
}
