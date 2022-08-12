package models

import (
	"errors"
	"regexp"
)

// ImageBlock is a subtype of block. It represents an image block.
type ImageBlock struct {
	block
	Image *ImageBlockObject `json:"image"`
}

// ImageBlockObject holds the detail of the ImageBlock.
type ImageBlockObject struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

// Validate performs validation to the ImageBlock. The supported extensions
// for field `Src` are jpg, jpeg, gif, and png.
func (s ImageBlock) Validate() (bool, []error) {
	// ImageBlock validation implementation
	var errs []error
	if s.Type != MBTImage {
		errs = append(errs, errors.New("invalid image block type"))
	}

	if match, _ := regexp.MatchString("^(http(s?):)([/|.|\\w|\\s|-])*\\.(?:jpg|jpeg|gif|png)$", s.Image.Src); !match {
		errs = append(errs, errors.New("invalid image src"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewImageBlock returns a new instance of a section block to be rendered
func NewImageBlock(src, alt string) *ImageBlock {
	block := ImageBlock{}
	block.Type = MBTImage
	block.Image.Src = src
	block.Image.Alt = alt

	return &block
}
