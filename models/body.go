package models

type MessageBodyType string

const (
	MBDTText   MessageBodyType = "text"
	MBDTImage  MessageBodyType = "image"
	MBDTButton MessageBodyType = "button"
)

type Body interface {
	BodyType() MessageBodyType
}

type Bodys struct {
	BodySet []Body `json:"body,omitempty"`
}

func NewBodys(bodys ...Body) Bodys {
	return Bodys{
		BodySet: bodys,
	}
}
