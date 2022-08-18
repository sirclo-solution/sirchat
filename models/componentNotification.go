package models

import (
	"errors"
)

type MessageNotificationObjectType string

const (
	// MNOTSuccess is the type for success notification
	MNOTSuccess MessageNotificationObjectType = "success"

	// MNOTFailed is the type for failed notification
	MNOTFailed MessageNotificationObjectType = "failed"
)

// NotificationComponent is a subtype of component. It represents a
// notfication component.
type NotificationComponent struct {
	component
	Notification NotificationObject `json:"notification"`
}

// NotificationObject holds the image detail for MessageObject.
type NotificationObject struct {
	// Type is the type of the notification.
	Type MessageNotificationObjectType `json:"type"`

	// Title is the text that will be shown as the title of notification.
	Title string `json:"title"`

	// Message is the text that will be shown as the body of notification.
	Message string `json:"message"`
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

// NewNotification used for initialization of new notification components
// NewNotification returns a new instance of a notification component to be rendered
func NewNotification(notificationComponentObj NotificationObject) *NotificationComponent {
	var c NotificationComponent
	c.Type = MCTNotification
	c.Notification = notificationComponentObj
	c.component.IComponent = &c
	return &c
}
