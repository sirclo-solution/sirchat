package models

import "errors"

// CarouselBlock defines a new block of carousel
type carouselBlock struct {
	block

	// Carousel contains the CarouselBlockObject that holds the detail of carousel block
	Carousel *CarouselBlockObject `json:"carousel,omitempty"`
}

// CarouselBlockObject defines detail of block carousel
type CarouselBlockObject struct {
	appendable

	// Title is the title of the carousel.
	// This field is required.
	Title string `json:"title"`

	// Descriptions is the array of descriptions for
	// each item in carousel block.
	// This field is required.
	Descriptions []string `json:"descriptions"`

	// Descriptions is the array of descriptions for
	// each item in carousel block.
	// This field is required.
	Images []imageBlock `json:"images"`
}

// Validate Carousel Block
func (ths *carouselBlock) Validate() (bool, []error) {
	var errs []error

	if len(ths.Carousel.Images) == 0 {
		errs = append(errs, errors.New("there must be at least one image in the carousel"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewCarouselBlock returns a new instance of a section block to be rendered
func NewCarouselBlock(title string) *carouselBlock {
	block := carouselBlock{
		Carousel: &CarouselBlockObject{
			Title: title,
		},
	}
	block.Type = MBTCarousel

	return &block
}

// AddDescriptionsCarousel for to adding descriptions field on carousel block
func (s *carouselBlock) AddDescriptionsCarousel(desc string) {
	s.Carousel.Descriptions = append(s.Carousel.Descriptions, desc)
}

// AddImageCarousel for to adding images field on carousel block
func (s *carouselBlock) AddImageCarousel(imgObj ImageBlockObject) {
	image := imageBlock{
		Image: &ImageBlockObject{
			Src:    imgObj.Src,
			Alt:    imgObj.Alt,
			Width:  imgObj.Width,
			Height: imgObj.Height,
		},
	}
	image.Type = MBTImage
	s.Carousel.Images = append(s.Carousel.Images, image)
}
