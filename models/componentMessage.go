package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MessageComponent struct {
	Component
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

func (ths *MessageComponent) Validate() (bool, error) {
	if ths.Type != MCTMessage {
		return false, errors.New("invalid message component type")
	}

	return true, nil
}

func (ths *MessageComponent) Compose() ([]byte, []error) {
	var errs []error

	if valid, err := ths.Validate(); !valid {
		err := fmt.Errorf("MessageComponent.Compose(): %s", err.Error())
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	res, err := json.Marshal(ths)
	if err != nil {
		return nil, []error{errors.New("error when marshaling message component")}
	}

	return res, nil
}

func NewMessage() *MessageComponent {
	var r MessageComponent
	r.Type = MCTMessage
	return &r
}
