package models

// ImageBody defines a new block of type section
type ImageBody struct {
	Type  MessageBodyType  `json:"type"`
	Image *ImageBodyObject `json:"image,omitempty"`
}

type ImageBodyObject struct {
	Alt string `json:"alt"`
	Src string `json:"src"`
}

// BlockType returns the type of the block
func (s ImageBody) BodyType() MessageBodyType {
	return s.Type
}

// NewImageBody returns a new instance of a section block to be rendered
func NewImageBody(imageObj *ImageBodyObject) *ImageBody {
	block := ImageBody{
		Type:  MBDTImage,
		Image: imageObj,
	}

	return &block
}
