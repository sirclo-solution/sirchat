package models

import (
	"errors"
	"regexp"
)

type ImageBlock struct {
	block
	Image *ImageBlockObject `json:"image"`
}

type ImageBlockObject struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

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
func NewImageBlock(imageBlockObj *ImageBlockObject) *ImageBlock {
	block := ImageBlock{
		Image: imageBlockObj,
	}
	block.Type = MBTImage

	return &block
}
