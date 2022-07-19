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
	BlockObject
	Src string `json:"src"`
	Alt string `json:"alt"`
}

func (s ImageBlock) Validate() error {
	// ImageBlock validation implementation
	if s.Type != "image" {
		return errors.New("invalid image block type")
	}

	if match, _ := regexp.MatchString("^(http(s?):)([/|.|\\w|\\s|-])*\\.(?:jpg|jpeg|gif|png)$", s.Image.Src); !match {
		return errors.New("invalid image src")
	}

	return nil
}

// NewImageBlock returns a new instance of a section block to be rendered
func NewImageBlock(imageBlockObj *ImageBlockObject) *ImageBlock {
	block := ImageBlock{
		Image: imageBlockObj,
	}
	block.Type = MBTImage

	return &block
}
