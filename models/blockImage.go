package models

import (
	"errors"
	"regexp"
)

type ImageBlock struct {
	Block
	Image *ImageBlockObject `json:"image"`
}

type ImageBlockObject struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

func (s ImageBlock) Validate() (bool, error) {
	// ImageBlock validation implementation
	if s.Type != MBTImage {
		return false, errors.New("invalid image block type")
	}

	if match, _ := regexp.MatchString("^(http(s?):)([/|.|\\w|\\s|-])*\\.(?:jpg|jpeg|gif|png)$", s.Image.Src); !match {
		return false, errors.New("invalid image src")
	}

	return true, nil
}

// NewImageBlock returns a new instance of a section block to be rendered
func NewImageBlock(imageBlockObj *ImageBlockObject) *ImageBlock {
	block := ImageBlock{
		Image: imageBlockObj,
	}
	block.Type = MBTImage

	return &block
}
