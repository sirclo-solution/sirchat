package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MessageNotificationObjectType string

const (
	MNOTSuccess MessageNotificationObjectType = "success"
	MNOTFailed  MessageNotificationObjectType = "failed"
)

type NotificationComponent struct {
	Component
	Notification NotificationObject `json:"notification"`
}

type NotificationObject struct {
	Type    MessageNotificationObjectType `json:"type"`
	Title   string                        `json:"title"`
	Message string                        `json:"message"`
}

func (ths *NotificationComponent) Validate() (bool, error) {
	if ths.Type != MCTNotification {
		return false, errors.New("invalid notification component type")
	}

	return true, nil
}

func (ths *NotificationComponent) Compose() ([]byte, []error) {
	var errs []error

	if valid, err := ths.Validate(); !valid {
		err := fmt.Errorf("NotificationComponent.Compose(): %s", err.Error())
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return nil, errs
	}

	res, err := json.Marshal(ths)
	if err != nil {
		return nil, []error{errors.New("error when marshaling notification component")}
	}

	return res, nil
}

func NewNotification() *NotificationComponent {
	var r NotificationComponent
	r.Type = MCTNotification
	return &r
}
