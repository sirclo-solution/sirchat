package models

type ImageBlock struct {
	Block
	Image *ImageBlockObject `json:"image"`
}

type ImageBlockObject struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

func (s ImageBlock) Validate() bool {
	// ImageBlock validation implementation

	return true
}

// NewImageBlock returns a new instance of a section block to be rendered
func NewImageBlock(imageBlockObj *ImageBlockObject) *ImageBlock {
	block := ImageBlock{
		Image: imageBlockObj,
	}
	block.Type = MBTImage

	return &block
}
