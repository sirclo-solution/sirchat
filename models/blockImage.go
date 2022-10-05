package models

import (
	"errors"
	"fmt"
	"regexp"
)

// ImageBlock is a subtype of block. It represents an image block.
type imageBlock struct {
	block

	// Image contains the ImageBlockObject that holds the detail of image block
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
func (ths *imageBlock) Validate() (bool, []error) {
	// ImageBlock validation implementation
	var errs []error
	if ths.Type != MBTImage {
		errs = append(errs, errors.New("invalid image block type"))
	}

	match, err := regexp.MatchString(`^(https:\/\/)([^\s(["<,>/]*[^/]*)(\/)[^\s[",><]*[^/]*((\?{1}|\#{1}).*)?$`, ths.Image.Src)
	if err != nil {
		errs = append(errs, fmt.Errorf("error when validating image src: %s", err.Error()))
	} else if !match {
		errs = append(errs, fmt.Errorf("invalid image src: %s", ths.Image.Src))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewImageBlock returns a new instance of a section block to be rendered
func NewImageBlock(imageObj ImageBlockObject) *imageBlock {
	block := imageBlock{}
	block.Type = MBTImage
	block.Image = &ImageBlockObject{
		Src:    imageObj.Src,
		Alt:    imageObj.Alt,
		Width:  imageObj.Width,
		Height: imageObj.Height,
	}

	return &block
}
