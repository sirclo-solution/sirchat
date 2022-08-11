package models

import (
	"errors"
)

type MessageNotificationObjectType string

const (
	MNOTSuccess MessageNotificationObjectType = "success"
	MNOTFailed  MessageNotificationObjectType = "failed"
)

// NotificationComponent is a subtype of component. It represents a
// notfication component.
type NotificationComponent struct {
	component
	Notification NotificationObject `json:"notification"`
}

// NotificationObject holds the image detail for MessageObject.
type NotificationObject struct {
	Type    MessageNotificationObjectType `json:"type"`
	Title   string                        `json:"title"`
	Message string                        `json:"message"`
}

// Validate performs validation to the NotificationComponent.
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

// NewNotification returns a new instance of a notification component to be rendered
func NewNotification(notificationType MessageNotificationObjectType, title, message string) *NotificationComponent {
	var c NotificationComponent
	c.Type = MCTNotification
	c.Notification.Type = notificationType
	c.Notification.Title = title
	c.Notification.Message = message
	c.component.IComponent = &c
	return &c
}
