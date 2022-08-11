package models

import (
	"errors"
)

// MessageComponent is a subtype of component. It represents a message component.
type MessageComponent struct {
	component
	Message MessageObject `json:"message"`
}

// MessageObject holds the detail of the MessageObject. The tenant, brand, and
// room ID for the message to be sent to are defined in this struct.
type MessageObject struct {
	TenantID string               `json:"tenant_id"`
	BrandID  string               `json:"brand_id"`
	RoomID   string               `json:"room_id"`
	Channel  string               `json:"channel"`
	Texts    []MessageTextObject  `json:"texts"`
	Images   []MessageImageObject `json:"images"`
}

// MessageTextObject holds the text body for MessageObject.
type MessageTextObject struct {
	Body string `json:"body"`
}

// MessageImageObject holds the image detail for MessageObject.
type MessageImageObject struct {
	Alt string `json:"alt"`
	Src string `json:"src"`
}

// Validate performs validation to the MessageComponent. The texts
// or images in the message component should be defined.
func (ths *MessageComponent) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MCTMessage {
		errs = append(errs, errors.New("invalid message component type"))
	}

	if ths.Message.TenantID == "" {
		errs = append(errs, errors.New("tenant ID in message component can't be empty"))
	}

	if ths.Message.BrandID == "" {
		errs = append(errs, errors.New("brand ID in message component can't be empty"))
	}

	if ths.Message.RoomID == "" {
		errs = append(errs, errors.New("room ID in message component can't be empty"))
	}

	if ths.Message.Channel == "" {
		errs = append(errs, errors.New("channel in message component can't be empty"))
	}

	if len(ths.Message.Texts) == 0 && len(ths.Message.Images) == 0 {
		errs = append(errs, errors.New("texts or images in message component can't be empty"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewMessage returns a new instance of a message component to be rendered
func NewMessage(tenantID, brandID, roomID, channel string) *MessageComponent {
	var c MessageComponent
	c.Type = MCTMessage
	c.Message.TenantID = tenantID
	c.Message.BrandID = brandID
	c.Message.RoomID = roomID
	c.Message.TenantID = tenantID
	c.component.IComponent = &c
	return &c
}

func NewMessageTextObject(body string) *MessageTextObject {
	return &MessageTextObject{
		Body: body,
	}
}

func NewMessageImageObject(src, alt string) *MessageImageObject {
	return &MessageImageObject{
		Src: src,
		Alt: alt,
	}
}
