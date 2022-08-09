package models

import (
	"errors"
)

type MessageNotificationObjectType string

const (
	MNOTSuccess MessageNotificationObjectType = "success"
	MNOTFailed  MessageNotificationObjectType = "failed"
)

type NotificationComponent struct {
	component
	Notification NotificationObject `json:"notification"`
}

type NotificationObject struct {
	Type    MessageNotificationObjectType `json:"type"`
	Title   string                        `json:"title"`
	Message string                        `json:"message"`
}

func (ths *NotificationComponent) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MCTNotification {
		errs = append(errs, errors.New("invalid notification component type"))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

func NewNotification() *NotificationComponent {
	var c NotificationComponent
	c.Type = MCTNotification
	c.component.IComponent = &c
	return &c
}
