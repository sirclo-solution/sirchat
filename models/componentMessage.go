package models

import (
	"errors"
)

// MessageComponent is a subtype of component. It represents a message component.
type messageComponent struct {
	component

	// Message contains the MessageObject that holds the detail of message component
	Message MessageObject `json:"message"`
}

// MessageObject holds the detail of the MessageObject. The tenant, brand, and
// room ID for the message to be sent to are defined in this struct.
type MessageObject struct {
	// TenantID contains tenant ID that the message will be sent from.
	// This field is required.
	TenantID string `json:"tenant_id"`

	// BrandID contains brand ID that the message will be sent from.
	// This field is required.
	BrandID string `json:"brand_id"`

	// RoomID contains the room ID that the message will be sent to.
	// This field is required.
	RoomID string `json:"room_id"`

	// Channel defines what channel the message will be sent through.
	// This field is required.
	Channel string `json:"channel"`

	// Texts defines the array of MessageTextObject that the message will contain.
	// This field is optional if the Images field is not empty.
	Texts []MessageTextObject `json:"texts"`

	// Images defines the array of MessageImageObject that the message will contain.
	// This field is optional if the Texts field is not empty.
	Images []MessageImageObject `json:"images"`
}

// MessageTextObject holds the text body for MessageObject.
type MessageTextObject struct {
	// Body contains the text body of the message.
	Body string `json:"body"`
}

// MessageImageObject holds the image detail for MessageObject.
type MessageImageObject struct {
	// Alt contains the alternative text to show when the image in Src is broken.
	// This field is required.
	Alt string `json:"alt"`

	// Src contains the URL of the image in the message.
	// This field is required.
	Src string `json:"src"`
}

// Validate performs validation to the MessageComponent. The texts
// or images in the message component should be defined.
func (ths *messageComponent) Validate() (bool, []error) {
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

// AddTextMessage adds MessageTextObject to MessageComponent.
func (ths *messageComponent) AddTextMessage(textBody string) {
	ths.Message.Texts = append(ths.Message.Texts, MessageTextObject{Body: textBody})
}

// AddImageMessage adds ImageTextObject to MessageComponent.
func (ths *messageComponent) AddImageMessage(alt, src string) {
	ths.Message.Images = append(ths.Message.Images, MessageImageObject{Alt: alt, Src: src})
}

// NewMessage used for initialization of new message components
// NewMessage returns a new instance of a message component to be rendered
func NewMessage(messageComponentObj MessageObject) *messageComponent {
	var c messageComponent
	c.Type = MCTMessage
	c.Message = messageComponentObj
	c.component.IComponent = &c
	return &c
}
