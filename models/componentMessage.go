package models

import (
	"errors"
)

type MessageComponent struct {
	component
	Message MessageObject `json:"message"`
}

type MessageObject struct {
	TenantID string              `json:"tenant_id"`
	BrandID  string              `json:"brand_id"`
	RoomID   string              `json:"room_id"`
	Channel  string              `json:"channel"`
	Texts    []MessageTextObject `json:"texts"`
}

type MessageTextObject struct {
	Body string `json:"body"`
}

func (ths *MessageComponent) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MCTMessage {
		errs = append(errs, errors.New("invalid message component type"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func NewMessage() *MessageComponent {
	var c MessageComponent
	c.Type = MCTMessage
	c.component.IComponent = &c
	return &c
}
