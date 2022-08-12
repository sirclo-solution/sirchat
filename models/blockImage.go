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
	// src is image url (jpg, jpeg, gif, png)
	Src string `json:"src"`

	// alt is text or label on image
	Alt string `json:"alt"`

	// Width of image (optional)
	// default 56px
	Width int `json:"width,omitempty"`

	// Width of image (optional)
	// default 56px
	Height int `json:"height,omitempty"`
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
func NewImageBlock(imageObj ImageBlockObject) *ImageBlock {
	block := ImageBlock{}
	block.Type = MBTImage
	block.Image = &ImageBlockObject{
		Src:    imageObj.Src,
		Alt:    imageObj.Alt,
		Width:  imageObj.Width,
		Height: imageObj.Height,
	}

	return &block
}
