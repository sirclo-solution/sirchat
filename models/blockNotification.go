package models

import "errors"

// NotificationBlock is a subtype of block. It represents a notification block.
type notificationBlock struct {
	block

	// Notification contains the NotificationBlockObject that holds the detail of notification block
	Notification *NotificationBlockObject `json:"notification"`
}

type NotificationBlockType string

const (
	// this type display information
	NotificationBlockTypeInfo NotificationBlockType = "info"

	// this type contains success information
	NotificationBlockTypeSuccess NotificationBlockType = "success"

	// this type contains waning information
	NotificationBlockTypeWarning NotificationBlockType = "warning"

	// this type contains error information
	NotificationBlockTypeError NotificationBlockType = "error"
)

// NotificationBlockObject holds the detail of the NotificationBlock.
type NotificationBlockObject struct {
	// Title of notification block (optional)
	Title string `json:"title"`

	// Message is the main content of the notification block
	// inform the reader of some pieces of information
	Message string `json:"message"`

	// Type determines what kind of information will be informed to the reader
	// the type also specifies what color of notification will be shown
	Type NotificationBlockType `json:"type"`
}

// Validate performs validation to the NotificationBlock. The supported value
// for field `Type` are success, info, warning, and error.
func (ths *notificationBlock) Validate() (bool, []error) {
	// NotificationBlock validation implementation
	var errs []error

	if ths.Type != MBTNotification {
		errs = append(errs, errors.New("invalid notification block type"))
	}

	if validNotificationBlockType := ths.Notification.Type.validateNotificationBlockType(); !validNotificationBlockType {
		errs = append(errs, errors.New("invalid type of notification, must be one of info|success|warning|error"))
	}

	if ths.Notification.Message == "" {
		errs = append(errs, errors.New("notification block body is required"))
	}

	return len(errs) == 0, errs
}

func (t NotificationBlockType) validateNotificationBlockType() bool {
	switch t {
	case NotificationBlockTypeInfo, NotificationBlockTypeSuccess, NotificationBlockTypeWarning, NotificationBlockTypeError:
		return true
	}

	return false
}

// NewNotificationBlock returns a new instance of a section block to be rendered
func NewNotificationBlock(notification NotificationBlockObject) *notificationBlock {
	block := new(notificationBlock)
	block.Type = MBTNotification
	block.Notification = &notification

	return block
}
