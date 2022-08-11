package models

import "errors"

// CarouselBlock defines a new block of carousel
type CarouselBlock struct {
	block
	Carousel *CarouselBlockObject `json:"carousel,omitempty"`
}

// CarouselBlockObject defines detail of block carousel
type CarouselBlockObject struct {
	appendable
	Title        string       `json:"title"`
	Descriptions []string     `json:"descriptions"`
	Images       []ImageBlock `json:"images"`
}

// Validate Carousel Block
func (s CarouselBlock) Validate() (bool, []error) {
	var errs []error

	if len(s.Carousel.Images) == 0 {
		errs = append(errs, errors.New("there must be at least one image in the carousel"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewCarouselBlock returns a new instance of a section block to be rendered
func NewCarouselBlock(title string) *CarouselBlock {
	block := CarouselBlock{
		Carousel: &CarouselBlockObject{
			Title: title,
		},
	}
	block.Type = MBTCarousel

	return &block
}

// AddDescriptionsCarousel for to adding descriptions field on carousel block
func (s *CarouselBlock) AddDescriptionsCarousel(desc string) {
	s.Carousel.Descriptions = append(s.Carousel.Descriptions, desc)
}

// AddImageCarousel for to adding images field on carousel block
func (s *CarouselBlock) AddImageCarousel(alt string, src string) {
	image := ImageBlock{
		Image: &ImageBlockObject{
			Src: src,
			Alt: alt,
		},
	}
	image.Type = MBTImage
	s.Carousel.Images = append(s.Carousel.Images, image)
}
